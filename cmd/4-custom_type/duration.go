package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
	"time"
)

// StarlarkDuration type for working with Durations in Starlark
type StarlarkDuration time.Duration

// String implement Value for StarlarkDuration
func (d StarlarkDuration) String() string {
	return fmt.Sprintf("%v", time.Duration(d))
}

// Type implement Value for StarlarkDuration
func (d StarlarkDuration) Type() string {
	return "duration"
}

// Freeze implement Value for StarlarkDuration
func (d StarlarkDuration) Freeze() {
}

// Truth implement Value for StarlarkDuration
func (d StarlarkDuration) Truth() starlark.Bool {
	if d == 0 {
		return starlark.False
	}
	return starlark.True
}

// Hash implement Value for StarlarkDuration
func (d StarlarkDuration) Hash() (uint32, error) {
	return uint32(d), nil
}

// Binary implement HasBinary in order to enable arithmetic operations on StarlarkDuration
func (d StarlarkDuration) Binary(op syntax.Token, y starlark.Value, side starlark.Side) (starlark.Value, error) {
	switch op {
	case syntax.PLUS:
		val, ok := y.(StarlarkDuration)
		if !ok {
			return nil, fmt.Errorf("invalid type:%t used for duration addition, only duration supported.",
				y)
		}
		return d + val, nil

	case syntax.MINUS:
		val, ok := y.(StarlarkDuration)
		if !ok {
			return nil, fmt.Errorf("invalid type:%t used for duration substraction, only duration supported.", y)
		}
		if side == starlark.Left {
			return d - val, nil
		} else {
			return val - d, nil
		}

	case syntax.SLASH:
		if side == starlark.Right {
			return nil, fmt.Errorf("a duration cannot appear to the right of a / operator")
		}
		intVal, ok := y.(starlark.Int)
		if ok {
			int64Val, ok := intVal.Int64()
			if !ok {
				return nil, fmt.Errorf("could not convert int to int64 %v", intVal)
			}
			return StarlarkDuration(int64(d) / int64Val), nil
		}
		floatVal, ok := y.(starlark.Float)
		if ok {
			f := float64(floatVal)
			return StarlarkDuration(time.Duration(int64(float64(d) / f))), nil
		}

		return nil, fmt.Errorf("unsupported type=%t for op /", y)

	case syntax.STAR:
		intVal, ok := y.(starlark.Int)
		if ok {
			int64Val, ok := intVal.Int64()
			if !ok {
				return nil, fmt.Errorf("could not convert int to int64 %v", intVal)
			}
			return StarlarkDuration(time.Duration(int64(d) * int64Val)), nil
		}
		floatVal, ok := y.(starlark.Float)
		if ok {
			f := float64(floatVal)
			return StarlarkDuration(time.Duration(int64(float64(d) * f))), nil
		}

		return nil, fmt.Errorf("unsupported type=%t for op *", y)
	default:
		return nil, nil // op not handled
	}
}

// Unary implement HasUnary in order to enable -/+ unary operations on StarlarkDuration
func (d StarlarkDuration) Unary(op syntax.Token) (starlark.Value, error) {
	switch op {
	case syntax.MINUS:
		return -d, nil
	case syntax.PLUS:
		return d, nil // nothing to do, support the + duration syntax
	default:
		return nil, nil
	}
}

// newDuration creates a starlark duration from a string, it is added as the `duration(s)` builtin
func newDuration(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var s string
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &s); err != nil {
		return nil, err
	}
	duration, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}
	return StarlarkDuration(duration), nil

}

// toDuration tries to convert a starlark Value to Duration, it supports strings ( which parses into a time.Duration) or
// StarlarkDuration which unwraps into a native time.Duration
func toDuration(val starlark.Value) (time.Duration, error) {
	switch val.(type) {
	case starlark.String:
		return time.ParseDuration(val.(starlark.String).GoString())
	case StarlarkDuration:
		return time.Duration(val.(StarlarkDuration)), nil
	}
	return 0, fmt.Errorf("cannot convert %T to duration", val)
}

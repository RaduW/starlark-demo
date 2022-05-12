package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"time"
)

func main() {

	var env = starlark.StringDict{
		"duration":    starlark.NewBuiltin("duration", newDuration),
		"Nanosecond":  StarlarkDuration(time.Nanosecond),
		"Microsecond": StarlarkDuration(time.Microsecond),
		"Millisecond": StarlarkDuration(time.Millisecond),
		"Second":      StarlarkDuration(time.Second),
		"Minute":      StarlarkDuration(time.Minute),
		"Hour":        StarlarkDuration(time.Hour),
		"Day":         StarlarkDuration(time.Hour * 24),
	}

	var thread = starlark.Thread{
		Name: "Starlark thread",
	}

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/4-custom_type/custom_type.py", nil, env)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

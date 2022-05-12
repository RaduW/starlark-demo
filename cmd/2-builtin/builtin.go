package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

// !!! HERE !!!
func doubleBuiltin(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (retVal starlark.Value, err error) {
	//NOTE !!! this only works for integers ( use startlark.Value and introspect if you want to make it generic)
	var val starlark.Int

	if err = starlark.UnpackArgs(b.Name(), args, kwargs, "val", &val); err != nil {
		return starlark.None, err
	}

	var param, ok = val.Int64()

	if !ok {
		return starlark.None, fmt.Errorf("int too large")
	}

	return starlark.MakeInt(int(2 * param)), nil
	// !!! END !!!
}

func main() {

	var env = starlark.StringDict{
		"GREET":  starlark.String("Welcome "),
		"TWO":    starlark.MakeInt(2),
		"double": starlark.NewBuiltin("double", doubleBuiltin),
	}

	var thread = starlark.Thread{
		Name: "Starlark thread",
	}

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/2-builtin/builtin.py", nil, env)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

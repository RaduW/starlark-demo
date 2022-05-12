package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

func doStuffBuiltin(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (retVal starlark.Value, err error) {
	//NOTE !!! this only works for integers ( use startlark.Value and introspect if you want to make it generic)
	var add int64 = 0
	var mul int64 = 1

	for _, kvRaw := range kwargs {
		k := kvRaw[0]
		v := kvRaw[1]
		if k == starlark.String("add") {
			add, _ = v.(starlark.Int).Int64()
		}
		if k == starlark.String("mul") {
			mul, _ = v.(starlark.Int).Int64()
		}
	}
	var resultRaw = []starlark.Value{}
	for _, valRaw := range args {
		val, _ := valRaw.(starlark.Int).Int64()
		newVal := val*mul + add
		resultRaw = append(resultRaw, starlark.MakeInt64(newVal))
	}
	return starlark.NewList(resultRaw), nil
}

func main() {

	var env = starlark.StringDict{
		"do_stuff": starlark.NewBuiltin("do_stuff", doStuffBuiltin),
	}

	var thread = starlark.Thread{
		Name: "Starlark thread",
	}

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/5-more_complex/more_complex.py", nil, env)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

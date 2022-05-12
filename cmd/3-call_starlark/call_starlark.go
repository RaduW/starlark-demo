package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

func main() {

	var thread = starlark.Thread{
		Name: "Starlark thread",
	}

	// NOTE: here we get the globals
	globals, err := starlark.ExecFile(&thread, "./cmd/3-call_starlark/call_starlark.py", nil, nil)

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
		return
	}

	///////////  say_hello BEGIN ////////
	var say_hello = globals["say_hello"]
	v, err := starlark.Call(&thread, say_hello, starlark.Tuple{starlark.String("Starlark")}, nil)

	if err != nil {
		fmt.Printf("Error calling say_hello")
		return
	}

	fmt.Printf("\n\n\n say_hello Returned: \n>>>%s<<<\n\n\n", v.String())
	///////////  say_hello END ////////

	///////////  add_all BEGIN ////////
	var add_all = globals["add_all"]
	var args = starlark.Tuple{}
	for i := 0; i <= 10; i++ {
		args = append(args, starlark.MakeInt(i))
	}

	v, _ = starlark.Call(&thread, add_all, args, nil)
	intVal, _ := toInt(v)
	fmt.Printf("\n\n\n add_all Returned: \n>>>%d<<<\n\n\n", intVal)
	///////////  say_hello END ////////

	///////////  add_all_vec BEGIN ////////
	var add_all_vec = globals["add_all_vec"]
	var lst_raw = []starlark.Value{}
	for i := 0; i <= 10; i++ {
		lst_raw = append(lst_raw, starlark.MakeInt(i))
	}
	var lst = starlark.NewList(lst_raw)
	v, _ = starlark.Call(&thread, add_all_vec, starlark.Tuple{lst}, nil)
	intVal, _ = toInt(v)
	fmt.Printf("\n\n\n add_al_vec Returned: \n>>>%d<<<\n\n\n", intVal)
	///////////  add_all_vec END ////////

}

func toInt(val starlark.Value) (int64, error) {
	intVal, ok := val.(starlark.Int)
	if !ok {
		return 0, fmt.Errorf("could not convert %v to int", val)
	}
	retVal, ok := intVal.Int64()
	if !ok {
		return 0, fmt.Errorf("could not convert %v to int64", intVal)
	}
	return retVal, nil
}

package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

func main() {

	// START

	var thread = starlark.Thread{}

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/1-hello/hello.py", nil, nil)
	fmt.Println("\n Script finished \n")

	// END

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

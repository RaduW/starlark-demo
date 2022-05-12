package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

var script = `
print("Embedded script")
def add(a,b):
	return a+b
print("1 + 2 = ", add(1,2))
`

func main() {

	// !!! HERE !!!
	var thread = starlark.Thread{
		Name: "Starlark thread",
	}
	// !!! END !!!

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "Fancy-File.py", script, nil)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

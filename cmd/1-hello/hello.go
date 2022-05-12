package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

// !!! HERE !!!
func CustomPrint(thread *starlark.Thread, msg string) {
	fmt.Println("Here's what they want to print:\n>>>")
	fmt.Print(msg)
	fmt.Println("\n<<<")
	// !!! END !!!
}

func main() {

	// !!! HERE !!!
	var thread = starlark.Thread{
		Name:  "Starlark thread",
		Print: CustomPrint,
	}
	// !!! END !!!

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/1-hello/hello.py", nil, nil)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

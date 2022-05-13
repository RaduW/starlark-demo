package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"strings"
)

var math = `
def add(a,b):
	return a+b

def subtract(a,b):
	return a-b
`

var greet = `
def say_hello(name):
	return "Hello %s" % name

def say_hi(name):
	return "Hi %s" % name
`

type moduleEntry struct {
	globals starlark.StringDict
	err     error
}

var embeddedModules = map[string]string{
	"math":  math,
	"greet": greet,
}

var moduleCache = make(map[string]*moduleEntry)

func loadModule(thread *starlark.Thread, moduleName string) (starlark.StringDict, error) {
	entry, ok := moduleCache[moduleName]
	if entry == nil {
		if ok {
			// we have set the module entry to nil so we have a circular dependency
			return nil, fmt.Errorf("cyclic dependency detected for module %s", moduleName)
		}
		// mark the fact that we are in the middle of resolving the module
		moduleCache[moduleName] = nil

		var globals starlark.StringDict
		var err error

		if isEmbedded(moduleName) {
			globals, err = loadEmbeddedModule(thread, moduleName)
		} else {
			globals, err = loadFileModule(thread, moduleName)
		}
		moduleCache[moduleName] = &moduleEntry{globals: globals, err: err}
		return globals, err
	}
	return entry.globals, entry.err
}

func isEmbedded(name string) bool {
	if _, ok := embeddedModules[name]; ok {
		return true
	}
	return false
}

func loadEmbeddedModule(thread *starlark.Thread, moduleName string) (starlark.StringDict, error) {
	var fileName = fmt.Sprintf("__builtin__/%s.py", moduleName)
	var script, _ = embeddedModules[moduleName]
	// NOTE starlark examples load in new threads ( don't know why) ?
	return starlark.ExecFile(thread, fileName, script, nil)
}

func loadFileModule(thread *starlark.Thread, moduleName string) (starlark.StringDict, error) {
	var fileName = strings.ReplaceAll(moduleName, ".", "/")
	fileName = fmt.Sprintf("./cmd/6-multi_module/%s.py", fileName)
	// NOTE starlark examples load in new threads ( don't know why) ?
	return starlark.ExecFile(thread, fileName, nil, nil)
}

func main() {

	var thread = starlark.Thread{
		Name: "Starlark thread",

		// !!! HERE !!!
		Load: loadModule,
		// !!! END !!!
	}

	fmt.Println("\n Running starlark script \n")
	var _, err = starlark.ExecFile(&thread, "./cmd/6-multi_module/main.py", nil, nil)
	fmt.Println("\n Script finished \n")

	if err != nil {
		fmt.Printf("ERROR:  Could not execute file\n%s", err)
	}
}

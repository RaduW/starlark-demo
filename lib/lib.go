package lib

import "go.starlark.net/starlark"

func CreateThread() *starlark.Thread {
	var thread = starlark.Thread{}

	return &thread
}

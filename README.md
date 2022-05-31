# starlark-demo
Some basic examples of embeding and using starlark-go.

## hello 
Demonstrates a minimal example, also how to provide your own print function

## builtin
Demonstrates how to add your go builtins (both constants and functions) that can then be called from a starlark scripts

## call_starlark
Demonstrates how to call functions defined in a starlark script from go.

## custom_type
Demonstrates how to provide custom types to starlark, it also demonstrates how to support binary (+,-,*,/) and unary (+/-) 
operations to custom types.

## more_complex
Demonstrates how to provide a builtin with a complex signature (i.e. *args, **kwargs ... )

## multi_module
Demonstrates how to support scripts with multiple modules, and shows an example of modules loaded from
both the file system and from strings embedded in the host app.

load("math", "add", "subtract")
load("greet", "say_hello", say_howdy="say_hi" )
load("x.y.math_plus", "multiply")

# Circular dependency below (will generate ERROR)
# load("circular1", "c1")

print("Loading 'main.py'")
print( "1 + 2 =", add(1,2))
print( "2 - 3 =", subtract(2,3))
print( say_hello("starlark"))
print( say_howdy("starlark"))
print( "3 * 4 =" , multiply(3,4))

# This is a Starlark file

def hello():
    v = 101
    x = "some-string"
    # NOTE ! go language string formatting
    print("""*** Hello from Starlark *** 
      the v=%d, x=%s""" % (v, x))


hello()

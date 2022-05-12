
def say_hello(name):
    return "Hello %s" %name

def add_all(*args):
    acc = 0
    for v in args:
        acc+=v
    return acc

def add_all_vec(v):
    return add_all(*v)
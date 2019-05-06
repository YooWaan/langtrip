def api_option(fn):
    def func_wrapper():
        return "wrap[{0}]".format(fn())
    return func_wrapper

@api_option
def say():
    return 'say'

def hello():
    ''' no decrated function '''
    return 'hello'

print(say())
print(hello())

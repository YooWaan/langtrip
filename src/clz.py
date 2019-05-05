class Parent:

    def say(self):
        print('py parent')

class Child(Parent):

    def say(self):
        print('py child')

Parent().say()
Child().say()



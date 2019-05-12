
def use_abc():
    from abc import ABCMeta, abstractmethod

    class Animal(metaclass=ABCMeta):

        @abstractmethod
        def run(self):
            pass

        def print(self, msg):
            print('abc: {}'.format(msg))

    class Dog(Animal):

        def run(self):
            self.print('wow')

    class Cat(Animal):

        def run(self):
            self.print('miao')

    class Runnable(metaclass=ABCMeta):

        @abstractmethod
        def run(self):
            pass

        def print(self, msg):
            print('run abc: {}'.format(msg))

    class Show(Runnable):

        def run(self):
            self.print('show')

    class Update(Runnable):

        def run(self):
            self.print('update')


    dog = Dog()
    cat = Cat()

    print('-- animal ---')
    dog.run()
    cat.run()

    dog = Show()
    cat = Update()
    print('-- runnable ---')
    dog.run()
    cat.run()

def not_use_abc():

    class Dog:

        def run(self):
            print('wow')

    class Cat:

        def run(self):
            print('miao')

    class Show:

        def run(self):
            print('show')

    class Update:

        def run(self):
            print('update')


    dog = Dog()
    cat = Cat()

    print('-- no abc animal ---')
    dog.run()
    cat.run()

    dog = Show()
    cat = Update()
    print('-- no abc runnable ---')
    dog.run()
    cat.run()


    
use_abc()
not_use_abc()

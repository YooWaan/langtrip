import java.util.function.Consumer;

interface Animal {
    void run();
}

class Dog implements Animal {
    @Override
    public void run() {
        System.out.println("wow");
    }
}

class Cat implements Animal {
    @Override
    public void run() {
        System.out.println("miao");
    }
}

class Show implements Runnable {
	public void run() {
		System.out.println("show");
	}
}

class Update implements Runnable {
	public void run() {
		System.out.println("update");
	}
}

Consumer<String> p = System.out::println;

p.accept("-- Animal interface --");
Animal dog = new Dog();
Animal cat = new Cat();
dog.run();
cat.run();

p.accept("-- Runnable interface --");
// cannot
p.accept("// Runnable dog2 = new Dog();");
Runnable show = new Show();
Runnable update = new Update();
show.run();
update.run();

/exit

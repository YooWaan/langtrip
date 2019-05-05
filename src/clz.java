class Parent {
    public void say() { System.out.println("java parent"); }
}
class Child extends Parent {
    public void say() { System.out.println("java child"); }
}

new Parent().say();
new Child().say();

/exit

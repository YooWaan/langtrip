package main

// Animal interface
type Animal interface {
	Run()
}

// Dog
type Dog struct{}
func (d Dog) Run() { println("wow") }

// Cat
type Cat struct{}
func (c Cat) Run() { println("miao") }

// Runnable interface
type Runnable interface {
	Run()
}

type Show struct{}
func (s Show) Run() { println("show")}
type Update struct{}
func (u Update) Run() { println("update")}

func main() {
	println("-- Animal --")
	var animal Animal
	animal = new(Dog)
	animal.Run()
	animal = new(Cat)
	animal.Run()

	println("-- Runnable --")
	var run Runnable
	run = new(Show)
	run.Run()
	run = new(Update)
	run.Run()

	println("-- Duck Typing --")
	run = new(Dog)
	run.Run()
	animal = new(Show)
	animal.Run()
}

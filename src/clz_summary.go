package main

type Say interface {
	Print()
}

type sayParent struct{}

func (p *sayParent) Print() { println("parent") }

type sayChild struct{}

func (p *sayChild) Print() { println("child") }

type Person struct {
	Say
}

func main() {
	p := Person{new(sayParent)}
	c := Person{new(sayChild)}

	println("-- parent --")
	p.Print()

	println("-- child --")
	c.Print()
}

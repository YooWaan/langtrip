package main

type Parent struct{}

func (p Parent) say() string { return "parent" }
func (p Parent) Print()      { println(p.say()) }
func (p Parent) Hello()      { println(p.say()) }

type Child struct{ Parent }

func (c Child) say() string { return "child" }
func (c Child) Hello()      { println(c.say()) }

func main() {
	p := Parent{}
	println("parent,", p.say())
	p.Print()
	p.Hello()

	c := Child{}
	println("child,", c.say())
	c.Print() // not work
	c.Hello()
}

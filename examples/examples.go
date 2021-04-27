package examples

import "fmt"

type Example interface {
	sayHi()
	sayBye()
	sayWhy()
}

type ExampleOne struct {
	thing1 string
	thing2 string
	thing3 string
}

func (example ExampleOne) sayHi() {
	fmt.Println("hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleOne) sayBye() {
	fmt.Println("bye")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleOne) sayWhy() {
	fmt.Println("why")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

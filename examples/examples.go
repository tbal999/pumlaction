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
	fmt.Println("example 1 says hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

type ExampleTwo struct {
	thing1 string
	thing2 string
	thing3 string
}

func (example ExampleTwo) sayHi() {
	fmt.Println("example 2 says hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func Hi(g geometry) {
	g.sayHi()
}


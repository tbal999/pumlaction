package examples

import "fmt"

type Example struct {
	thing1 string
	thing2 string
	thing3 string
}

func (example Example) sayHi() {
	fmt.Println("hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example Example) sayBye() {
	fmt.Println("bye")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example Example) sayWhy() {
	fmt.Println("why")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

type ExampleTwo struct {
	thing1 string
	thing2 string
	thing3 string
}

func (example ExampleTwo) sayHi() {
	fmt.Println("hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleTwo) sayBye() {
	fmt.Println("bye")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleTwo) sayWhy() {
	fmt.Println("why")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

type ExampleThree struct {
	thing1 string
	thing2 string
	thing3 string
}

func (example ExampleThree) sayHi() {
	fmt.Println("hello")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleThree) sayBye() {
	fmt.Println("bye")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleThree) sayWhy() {
	fmt.Println("why")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

func (example ExampleThree) sayAgain() {
	fmt.Println("again?")
	fmt.Println(example.thing1, example.thing2, example.thing3)
}

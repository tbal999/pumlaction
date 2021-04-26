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

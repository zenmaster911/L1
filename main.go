package main

import "fmt"

type Human struct {
}

func (h Human) HelloWorld() {
	fmt.Println("Hello world!")
}

func (h Human) OtherMethod() {
	fmt.Println("Other Method")
}

type Action struct {
	Human
}

func main() {
	var Say Action
	Say.HelloWorld()

}

package main

import (
	"./security"
	"./something"
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) Welcome() string {
	return fmt.Sprintf("Welcome %v! You are %v.", p.name, p.age)
}

func main() {
	me := Person{name: "George", age: 26}

	fmt.Println(me.Welcome())
	something.A()
	something.AAA()

	secure_word, err := security.Hash("My word")
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Printf("Hash is: %x\n", secure_word)
	}
}

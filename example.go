package main

import (
	"fmt"
	"github.com/georgedrummond/learning-go/something"
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
}

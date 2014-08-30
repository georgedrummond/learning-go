package main

import (
	"./gravatar"
	"./pages"
	"./security"
	"./something"
	"fmt"
  "encoding/json"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Welcome() string {
	return fmt.Sprintf("Welcome %v! You are %v.", p.Name, p.Age)
}

func (p *Person) ToJson() ([]byte, error) {
	json, error := json.Marshal(p)

  fmt.Println(string(json))

  return json, error
}

func main() {
	me := Person{Name: "George", Age: 26}

	fmt.Println(me.Welcome())
	something.A()
	something.AAA()

	secure_word := security.Hash("My word")
	fmt.Printf("Hash is: %x\n", secure_word)

	page := pages.Page()
	fmt.Println(page)

	gravatar_url := gravatar.EmailUrl("drummond@rentify.com")
	fmt.Println(gravatar_url)

  json, err := me.ToJson()

  if err == nil {
    fmt.Println(string(json))
  }
}

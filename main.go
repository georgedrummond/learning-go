package main

import (
	"encoding/json"
	"fmt"
	"github.com/georgedrummond/learning-go/gravatar"
	"github.com/georgedrummond/learning-go/pages"
	"github.com/georgedrummond/learning-go/security"
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

	secure_word := security.Hash("My word")
	fmt.Printf("Hash is: %x\n", secure_word)

	page := pages.Page()
	fmt.Println(page)

	profile := gravatar.Profile{Email: "drummond@rentify.com"}
	fmt.Println(profile.Url())

	json, err := me.ToJson()

	if err == nil {
		fmt.Println(string(json))
	}

	security.Blah()
}

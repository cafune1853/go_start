package go_start

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}


func (p *Person) String() string {
	fmt.Println()
	return p.Name
}

func (p *Person) Speak(word string) string {
	return p.Name + " saying " + word
}

type Speaker interface {
	Speak(word string) string
}

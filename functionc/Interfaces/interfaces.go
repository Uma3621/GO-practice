package interfaces

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Pikachu struct {
}

func (p Pikachu) Speak() string {
	return "Pika pika!"
}

type Programmer struct {
}

func (j Programmer) Speak() string {
	return "Design patterns!"
}
func main() {
	animals := []Animal{Dog{}, Cat{}, Pikachu{}, Programmer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (dog *Dog) speak() string {
	return "Woof"
}

type Cat struct {
}

func (cat *Cat) speak() string {
	return "Meow"
}

func AnimalFactory(animal string) Animal {
	switch animal {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}

func main() {
	// Create a dog
	dog := AnimalFactory("dog")
	fmt.Println(dog.speak()) // Output: Woof

	// Create a cat
	cat := AnimalFactory("cat")

	fmt.Println(cat.speak()) // Output: Meow
}

package main

import "fmt"

type MyInterface interface {
	// 定义接口方法
	SayHello()
}

type Person struct {
	Name string
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

type Dog struct {
	Name string
}

func (d Dog) SayHello() {
	fmt.Println("Hello, my name is", d.Name)
}

func say(aa MyInterface) (a string) {
	aa.SayHello()
	return "aaa"
}

func main() {
	fmt.Println("main started")
	var obj = Person{Name: "Alice"}
	var obj1 = Dog{Name: "Bob"}
	var string1 = say(obj)
	var string2 = say(obj1)
	fmt.Println(string1)
	fmt.Println(string2)
}

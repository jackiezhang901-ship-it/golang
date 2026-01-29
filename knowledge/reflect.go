package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// 获取变量的反射值
	v := reflect.ValueOf(p)

	// 通过反射获取方法
	method := v.MethodByName("Greet")

	// 检查方法是否存在且可调用
	if method.IsValid() {
		// 调用方法，参数为 nil，因为 Greet 没有入参
		method.Call(nil)
	}
}

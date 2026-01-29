package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance 获取单例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1 == s2)
}

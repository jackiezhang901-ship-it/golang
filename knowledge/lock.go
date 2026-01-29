package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start to")
	wg.Add(2)
	go func() {
		fmt.Println("goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("goroutine 2")
		wg.Done()
	}()

	wg.Wait()

}

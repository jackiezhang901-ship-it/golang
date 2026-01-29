package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	slices := make([]int, 0, 10)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	slices = append(slices, 1)
	fmt.Println(slices)
	maps := make(map[string]int)
	maps["a"] = 1
	maps["b"] = 2
	fmt.Println(maps)

	ch := make(chan int, 3)

	// 发送数据（不需要立即有接收者）
	ch <- 1
	ch <- 2
	ch <- 3

	// 接收数据
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
}

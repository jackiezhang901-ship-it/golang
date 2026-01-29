package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// 在函数退出时通知 WaitGroup，当前任务已完成
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// 启动 5 个 worker goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 为每个 worker 增加一个计数
		go worker(i, &wg)
	}

	// 等待所有 worker 都完成
	wg.Wait()
	fmt.Println("All workers finished.")
}

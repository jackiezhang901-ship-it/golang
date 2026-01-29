package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建可取消的Context
	ctx, cancel := context.WithCancel(context.Background())

	// 创建子Context
	childCtx, childCancel := context.WithCancel(ctx)

	go func() {
		select {
		case <-childCtx.Done():
			fmt.Println("子Context被取消:", childCtx.Err())
		case <-ctx.Done():
			fmt.Println("父Context被取消:", ctx.Err())
		}
	}()

	// 取消父Context，会级联取消子Context
	cancel()
	time.Sleep(1 * time.Second)

	// 注意：即使childCancel没有被调用，子Context也会被取消
	childCancel() // 不会造成问题，可以安全调用
}

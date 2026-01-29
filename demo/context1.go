package main

import (
	"context"
	"fmt"
)

type contextKey string

const userIDKey contextKey = "userID"

func processRequest(ctx context.Context) {
	// 获取值
	if userID := ctx.Value(userIDKey); userID != nil {
		fmt.Printf("处理用户 %s 的请求\n", userID)
	}
}

func main() {
	// 创建带值的Context
	ctx := context.WithValue(context.Background(), userIDKey, "12345")
	processRequest(ctx)
}

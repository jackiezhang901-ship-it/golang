package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 设置超时
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 模拟耗时操作
	select {
	case <-time.After(8 * time.Second):
		fmt.Fprintln(w, "请求处理成功")
	case <-ctx.Done():
		http.Error(w, "请求超时", http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

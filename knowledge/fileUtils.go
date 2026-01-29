package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件

	file, err := os.Open("D:\\download\\huawei_cloud.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("读取文件失败:", err)
	}
	fmt.Println("文件内容:", string(content[:n]))
}

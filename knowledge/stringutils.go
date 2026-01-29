package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(path.Join("dsa", "cc"))
}

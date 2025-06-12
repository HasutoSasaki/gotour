package main

import "fmt"

// 関数の定義と使用
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

package main

import "fmt"

// 複数の戻り値を持つ関数
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

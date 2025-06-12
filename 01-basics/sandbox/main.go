package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 乱数の生成
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}

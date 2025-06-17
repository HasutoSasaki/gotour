package main

import "fmt"

func main() {
	names := [4]string{"Alice", "Bob", "Charlie", "Diana"}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	a[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

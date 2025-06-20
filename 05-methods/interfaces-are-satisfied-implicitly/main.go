package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements interface I.
// but we don't need to explicitly declare that it does so.
func (f T) M() {
	fmt.Println(f.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

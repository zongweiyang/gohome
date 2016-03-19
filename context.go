package main

import "fmt"

var a int

func main() {
	a = 5
	fmt.Println(a)
	q()
	p()
}

func p() {
	fmt.Println(a)
}
func q() {
	a := 6
	fmt.Println(a)
	p()
}

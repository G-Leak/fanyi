package main

import "fmt"

func main() {
	a := 13
	b := 23

	c := compare(a, b)
	fmt.Println(c)
	fmt.Println("hello world!")
}

func compare(a int, b int) int {
	return a + b
}

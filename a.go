package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	a := 13
	b := 23

	c := compare(a, b)
	fmt.Println(c)
	fmt.Println("hello world!")

	arr := []float64{1, 2, 5, 6}
	v := stat.Variance(arr, nil)
	fmt.Printf("方差=%f", v)
}

func compare(a int, b int) int {
	return a + b
}

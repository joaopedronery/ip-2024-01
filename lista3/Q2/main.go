package main

import "fmt"

func main() {

	var n float64
	var k int

	fmt.Scan(&n)

	for n < 1{
		fmt.Scan(&n)
	}

	x := make([]int, int(n), 1000)

	for i := 0; i < int(n); i++{
		fmt.Scan(&x[i])
	}

	fmt.Scan(&k)

	var y []int

	for _, v := range x{
		if v >= k{
			y = append(y, v)
		}
	}
	fmt.Print(len(y))
}
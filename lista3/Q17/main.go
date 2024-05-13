package main

import(
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 4999)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	contar := make(map[int]int)

	for _, num := range x{
		contar[num]++
	}

	y := make([]int, 0)

	for _, v := range contar{
		if v == 1{
			y = append(y, v)
		}
	}
	fmt.Print(len(y))
}
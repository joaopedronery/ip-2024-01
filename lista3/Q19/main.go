package main

import(
	"fmt"
)

func main(){

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 1000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	contar := make(map[int]int)

	for _, num := range x{
		contar[num]++
	}

	for i, _ := range contar{
		fmt.Printf("%d\n", i)
	}
}
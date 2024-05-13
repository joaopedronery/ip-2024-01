package main

import (
	"fmt"
	"sort"
)
func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 1000000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	sort.Ints(x)

	var mediana float64

	if n%2 == 0{
		mediana = float64(x[n/2-1] + x[n/2]) / 2
	} else {
		mediana = float64(x[n/2])
	}
	fmt.Printf("%.2f\n", mediana)
}
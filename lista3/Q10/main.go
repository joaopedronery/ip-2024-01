package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 10000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	ui := len(x) - 1
	un := x[ui]

	var y []int
	
	for i := 0; i < n; i++{
		if un == x[i]{
			y = append(y, x[i])
		}
	}
	fmt.Printf("Nota %d, %d vezes\n", un, len(y))

	mn := -1
	ind := -1
	for i := 0; i < n; i++{
		if x[i] > mn {
			mn = x[i]
			ind = i
		}
	}
	fmt.Printf("Nota %d, indice %d\n", mn, ind)
}

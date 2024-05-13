package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 4999)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	soma := 0

	for i := 0; i < n; i++{
			soma += x[i]
	}
	fmt.Print(soma)
}
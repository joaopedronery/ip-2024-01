package main

import "fmt"

func main() {

	var n, ni int

	fmt.Scan(&n)

	x := make([]int, n, 100000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	fmt.Scan(&ni)

	y := make([]int, ni, 1000)

	for i := 0; i < ni; i++{
		fmt.Scan(&y[i])
	}

	for i := 0; i < ni; i++{
	var encontrado bool

	for _, v := range x{
		if y[i] == v{
			encontrado = true
		}
		}
	if encontrado{
		fmt.Print("ACHEI\n")
	} else{
		fmt.Print("NAO ACHEI\n")
	}
}
}
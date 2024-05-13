package main

import (
	"fmt"
)

func main() {
	var n, k, alunop int
	fmt.Scan(&n, &k)

	chegada := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&chegada[i])
		if chegada[i] <= 0 {
			alunop++
		}
	}

	if alunop < k {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")

		for i := n - 1; i >= 0; i-- {
			if chegada[i] <= 0 {
				fmt.Printf("%d\n", i+1)
			}
		}
	}
}
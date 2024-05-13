package main

import (
	"fmt"
)

func main() {

	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		x := make([]int, 11)
		for j := range x {
			fmt.Scan(&x[j])
		}
		somam := (x[0]*1 + x[1]*2 + x[2]*3 + x[3]*4 + x[4]*5 + x[5]*6 + x[6]*7 + x[7]*8 + x[8]*9) % 11
		somaim := (x[8]*1 + x[7]*2 + x[6]*3 + x[5]*4 + x[4]*5 + x[3]*6 + x[2]*7 + x[1]*8 + x[0]*9) % 11

		if somam == 10 {
			somam = 0
		}
		if somaim == 10 {
			somaim = 0
		}

		if x[9] == somam && x[10] == somaim {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}
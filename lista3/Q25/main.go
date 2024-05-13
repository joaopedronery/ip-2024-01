package main

import(
	"fmt"
)

func main() {
	
	var sorteados [6]int

	for i := 0; i < 6; i++{
	fmt.Scan(&sorteados[i])
	}

	var n int
	fmt.Scan(&n)

	quadra, quina, sena := 0, 0, 0

	for i := 0; i < n; i++{
		var aposta [6]int
		for j := 0; j < 6; j++{
			fmt.Scan(&aposta[j])
		}

		acertos := 0

		for _, sorteado := range sorteados{
			for _, apostado := range aposta{
				if sorteado == apostado{
					acertos++
				}
			}
		}

		switch acertos {
		case 6:
			sena++
		case 5:
			quina++
		case 4:
			quadra++
		}
	}

	if sena == 0{
		fmt.Println("Nao houve acertador para sena")
	} else {
		fmt.Printf("Houve %d acertador(es) da sena\n", sena)
	}

	if quina == 0{
		fmt.Println("Nao houve acertador para quina")
	} else {
		fmt.Printf("Houve %d acertador(es) da quina\n", quina)
	}

	if quadra == 0{
		fmt.Println("Nao houve acertador para quadra")
	} else {
		fmt.Printf("Houve %d acertador(es) da quadra\n", quadra)
	}
}
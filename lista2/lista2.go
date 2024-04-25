package main

import (
	"fmt"
	"math"
)

func main() {
	questao21();
}

func questao1() {
	var messages []string;
	var matricula, frequencia int;
	var prova1, prova2, prova3, prova4, prova5, prova6, prova7, prova8, lista1, lista2, lista3, lista4, lista5, trabalhoFinal float64;

	for {
		var finalSituationMessage, fullMessage string;
		fmt.Scanln(&matricula, &prova1, &prova2, &prova3, &prova4, &prova5, &prova6, &prova7, &prova8, &lista1, &lista2, &lista3, &lista4, &lista5, &trabalhoFinal, &frequencia);
		mediaProvas := (prova1 + prova2 + prova3 + prova4 + prova5 + prova6 + prova7 + prova8) / 8;
		mediaLista := (lista1 + lista2 + lista3 + lista4 + lista5) / 5;
		notaFinal := 0.7 * mediaProvas + 0.15 * mediaLista + 0.15 * trabalhoFinal;
		
		if mediaProvas == -1 && mediaLista == -1 && matricula == -1 && frequencia == -1 && trabalhoFinal == -1 {
			break
		}

		if float64(frequencia) / 128 < 0.75 {
			if notaFinal < 6 {
				finalSituationMessage = "REPROVADO POR NOTA E POR FREQUENCIA";
			} else {
				finalSituationMessage = "REPROVADO POR FREQUENCIA";
			}
		} else {
			if notaFinal < 6 {
				finalSituationMessage = "REPROVADO POR NOTA";
			} else {
				finalSituationMessage = "APROVADO";
			}
		}

		fullMessage = fmt.Sprintf("Matricula: %d, Nota Final: %.2f, Situacao Final: %s", matricula, notaFinal, finalSituationMessage);
		messages = append(messages, fullMessage);
	}

	for i := range messages {
		fmt.Println(messages[i]);
	}
}

func questao2() {
	var populacaoA, populacaoB, anos int;
	var novaPopulacaoA, novaPopulacaoB float64;

	fmt.Scan(&populacaoA, &populacaoB);

	novaPopulacaoA = float64(populacaoA);
	novaPopulacaoB = float64(populacaoB);

	for novaPopulacaoA < novaPopulacaoB {
		anos = anos + 1;
		novaPopulacaoA = novaPopulacaoA * 1.03;
		novaPopulacaoB = novaPopulacaoB * 1.015;
	}

	fmt.Printf("ANOS = %d\n", anos);
}

func questao3() {
	var numero, somatorio int;
	
	fmt.Scan(&numero);
	somatorio = numero

	if numero == 0 {
		fmt.Println("0! = 1");
		return
	}

	for i := numero - 1; i >= 1; i-- {
		somatorio = somatorio * i;
	}

	fmt.Printf("%d! = %d\n", numero, somatorio);
}

func questao4() {
    var n, i, K, s float64

    fmt.Println("Digite um número (n de 0 a 9):")
    fmt.Scanln(&n)
    fmt.Println("Digite o valor inicial (i):")
    fmt.Scanln(&i)
    fmt.Println("Digite a quantidade de valores (K):")
    fmt.Scanln(&K)
    fmt.Println("Digite o incremento (s):")
    fmt.Scanln(&s)

    fmt.Println("Tabuada de soma:")
    for j := 0.0; j < K; j++ {
        B := i + j*s
        resultado := n + B
        fmt.Printf("%.2f + %.2f = %.2f\n", n, B, resultado)
    }

    fmt.Println("Tabuada de subtração:")
    for j := 0.0; j < K; j++ {
        B := i + j*s
        resultado := n - B
        fmt.Printf("%.2f - %.2f = %.2f\n", n, B, resultado)
    }

    fmt.Println("Tabuada de multiplicação:")
    for j := 0.0; j < K; j++ {
        B := i + j*s
        resultado := n * B
        fmt.Printf("%.2f * %.2f = %.2f\n", n, B, resultado)
    }

    fmt.Println("Tabuada de divisão:")
    for j := 0.0; j < K; j++ {
        B := i + j*s
        resultado := n / B
        fmt.Printf("%.2f / %.2f = %.2f\n", n, B, resultado)
    }
}

func questao5() {
    var n int
    fmt.Println("Digite o número de elementos da sequência:")
    fmt.Scanln(&n)

    fmt.Println("Digite a sequência de números inteiros:")
    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }

    comprimentoMaximo := 0
    comprimentoAtual := 1
    for i := 1; i < n; i++ {
        if numeros[i] > numeros[i-1] {
            comprimentoAtual++
        } else {
            if comprimentoAtual > comprimentoMaximo {
                comprimentoMaximo = comprimentoAtual
            }
            comprimentoAtual = 1
        }
    }
    if comprimentoAtual > comprimentoMaximo {
        comprimentoMaximo = comprimentoAtual
    }

    fmt.Printf("O comprimento do segmento crescente máximo é: %d\n", comprimentoMaximo-1)
}

func questao6() {
    var n int
    fmt.Println("Digite o número de elementos da sequência:")
    fmt.Scanln(&n)

    fmt.Println("Digite a sequência de números inteiros:")
    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }

    maiorSubsequencia := 1
    subsequenciaAtual := 1
    for i := 1; i < n; i++ {
        if numeros[i] == numeros[i-1] {
            subsequenciaAtual++
        } else {
            if subsequenciaAtual > maiorSubsequencia {
                maiorSubsequencia = subsequenciaAtual
            }
            subsequenciaAtual = 1
        }
    }
    if subsequenciaAtual > maiorSubsequencia {
        maiorSubsequencia = subsequenciaAtual
    }

    fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maiorSubsequencia)
}

func questao7() {
    var (
        num, somaPar, somaImpar float64
        countPar, countImpar float64
    )

    fmt.Println("Digite uma sequência de números inteiros (digite 0 para encerrar):")
    for {
        fmt.Scan(&num)
        if num == 0 {
            break
        }
        if num % 2 == 0 {
            somaPar += num
            countPar++
        } else {
            somaImpar += num
            countImpar++
        }
    }

    mediaPar := somaPar / countPar
    mediaImpar := somaImpar / countImpar

    fmt.Printf("MEDIA PAR = %.2f\n", mediaPar)
    fmt.Printf("MEDIA IMPAR = %.2f\n", mediaImpar)
}

func questao8() {
    var N int

    fmt.Println("Digite a quantidade de times no campeonato:")
    fmt.Scanln(&N)

    if N < 2 {
        fmt.Println("Campeonato inválido!")
        return
    }

    contadorFinal := 1
    for i := 1; i <= N; i++ {
        for j := i + 1; j <= N; j++ {
            fmt.Printf("Final %d: Time%d X Time%d\n", contadorFinal, i, j)
            contadorFinal++
        }
    }
}

func questao9() {
    var N int

    fmt.Println("Digite um número inteiro positivo:")
    _, err := fmt.Scanln(&N)
    if err != nil || N <= 0 {
        fmt.Println("Número inválido!")
        return
    }

    primo := true
    for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
        if N%i == 0 {
            primo = false
            break
        }
    }

    if primo && N != 1 {
        fmt.Println("PRIMO")
    } else {
        fmt.Println("NAO PRIMO")
    }
}

func questao10() {
    for {
        var matricula int
        var horasTrabalhadas, valorPorHora float64

        _, err := fmt.Scan(&matricula, &horasTrabalhadas, &valorPorHora)
        if err != nil {
            fmt.Println("Erro ao ler entrada:", err)
            return
        }

        if matricula == 0 {
            break
        }

        salario := horasTrabalhadas * valorPorHora

        fmt.Printf("%d %.2f\n", matricula, salario)
    }
}


func questao11() {
    var n, e float64

    fmt.Println("Digite o número cuja raiz quadrada deseja-se obter:")
    fmt.Scanln(&n)
    fmt.Println("Digite o erro desejado:")
    fmt.Scanln(&e)

    r := 1.0

    for {
        r_antigo := r
        r = (r_antigo + n/r_antigo) / 2

        erro := math.Abs(n - (r * r))

        fmt.Printf("%.9f %.9f\n", r, erro)

        if erro <= e {
            break
        }
    }
}


func questao12() {
    var ValorIngresso, ValorInicial, ValorFinal float64

    fmt.Println("Digite o valor do ingresso:")
    fmt.Scanln(&ValorIngresso)
    fmt.Println("Digite o valor inicial do intervalo:")
    fmt.Scanln(&ValorInicial)
    fmt.Println("Digite o valor final do intervalo:")
    fmt.Scanln(&ValorFinal)

    if ValorInicial >= ValorFinal {
        fmt.Println("INTERVALO INVALIDO.")
        return
    }

    melhorLucro := 0.0
    melhorValor := 0.0
    melhorNumeroIngressos := 0

    for preco := ValorInicial; preco <= ValorFinal; preco += 1.0 {
        numeroIngressos := calcularIngressos(preco)

        lucro := calcularLucro(ValorIngresso, preco, numeroIngressos)

        if lucro > melhorLucro {
            melhorLucro = lucro
            melhorValor = preco
            melhorNumeroIngressos = numeroIngressos
        }

        fmt.Printf("V: %.2f, N: %d, L: %.2f\n", preco, numeroIngressos, lucro)
    }

    fmt.Println("\nResumo:")
    fmt.Printf("Melhor valor final: %.2f\n", melhorValor)
    fmt.Printf("Lucro: %.2f\n", melhorLucro)
    fmt.Printf("Numero de ingressos: %d\n", melhorNumeroIngressos)
}

func calcularIngressos(preco float64) int {
    var numeroIngressos int

    if preco >= 1.0 {
        numeroIngressos = 120 - int(preco-1.0)*25
    } else {
        numeroIngressos = 120 + int(1.0-preco)*30
    }

    return numeroIngressos
}

func calcularLucro(ValorIngresso, preco float64, numeroIngressos int) float64 {
    despesasFixas := 200.0 + float64(numeroIngressos)*0.05
    receita := preco * float64(numeroIngressos)
    lucro := receita - ValorIngresso*float64(numeroIngressos) - despesasFixas
    return lucro
}

func questao13() {
	var n int

	
	fmt.Println("Digite um número inteiro:")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}

	
	totalGrains := uint64(0)
	grains := uint64(n)

	
	for linha := 0; linha < 8; linha++ {
		for coluna := 0; coluna < 8; coluna++ {
	
			if (linha+coluna)%2 == 0 {
				totalGrains += grains
			}
	
			grains *= 2
		}
	}

	
	fmt.Println(totalGrains)
}


func questao14() {
    var m, n int

    
    fmt.Println("Digite o número de linhas (m) e colunas (n) da matriz:")
    fmt.Scanln(&m, &n)

    
    if m <= 0 || n <= 0 {
        fmt.Println("As dimensões da matriz devem ser positivas.")
        return
    }

    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
    
            if i < j {
    
                fmt.Printf("(%d-%d) ", i, j)
            }
        }
    
        fmt.Println()
    }
}

func questao15() {
	var T int

	
	fmt.Scanln(&T)

	
	for i := 0; i < T; i++ {
		var A, B int

	
		fmt.Scanln(&A, &B)

	
		A = inverterNumero(A)
		B = inverterNumero(B)

	
		maiorNumero := A
		if B > A {
			maiorNumero = B
		}

	
		fmt.Println(maiorNumero)
	}
}


func inverterNumero(num int) int {

	d1 := num / 100
	d2 := (num % 100) / 10
	d3 := num % 10


	invertido := d3*100 + d2*10 + d1

	return invertido
}

func questao16() {
	var n int

	
	fmt.Println("Digite um número inteiro maior que zero:")
	fmt.Scanln(&n)

	
	if n <= 0 {
		fmt.Println("O número deve ser maior que zero.")
		return
	}

	
	for hip := 1; hip <= n; hip++ {
	
		for c1 := 1; c1 < hip; c1++ {
			for c2 := c1; c2 < hip; c2++ {
	
				if c1*c1+c2*c2 == hip*hip {
	
					fmt.Printf("hipotenusa = %d, catetos %d e %d\n", hip, c1, c2)
				}
			}
		}
	}
}

func questao17() {
	var codigo, vendas int
	var precoCompra, precoVenda, lucroTotal, lucro, maiorLucro, maiorLucroCodigo float64
	var mercadoriasMenor10, mercadoriasEntre10e20, mercadoriasMaior20 int
	var maisVendido, maisVendidoCodigo int
	var totalCompra, totalVenda float64

	for {
		
		_, err := fmt.Scanln(&codigo, &precoCompra, &precoVenda, &vendas)
		if err != nil {
			break 
		}

		
		lucro = (precoVenda - precoCompra) * float64(vendas)
		lucroTotal += lucro

		
		if lucro < precoCompra*0.1 {
			mercadoriasMenor10++
		} else if lucro >= precoCompra*0.1 && lucro <= precoCompra*0.2 {
			mercadoriasEntre10e20++
		} else {
			mercadoriasMaior20++
		}

		
		if lucro > maiorLucro {
			maiorLucro = lucro
			maiorLucroCodigo = float64(codigo)
		}

		
		if vendas > maisVendido {
			maisVendido = vendas
			maisVendidoCodigo = codigo
		}

		
		totalCompra += precoCompra * float64(vendas)
		totalVenda += precoVenda * float64(vendas)
	}

	
	percentualLucro := (lucroTotal / totalCompra) * 100

	
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", mercadoriasMenor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", mercadoriasEntre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", mercadoriasMaior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %.0f\n", maiorLucroCodigo)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", maisVendidoCodigo)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompra, totalVenda, percentualLucro)
}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}


func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}

func questao18() {
	var a, b, c int


	fmt.Println("Digite três números inteiros diferentes de zero:")
	fmt.Scanln(&a, &b, &c)


	mmcAB := mmc(a, b)
	mmcTotal := mmc(mmcAB, c)


	fmt.Printf("%5d %5d %5d :%d\n", a, b, c, mmcTotal)
}

func questao19() {
	var m int

	
	fmt.Println("Digite um número inteiro maior que zero:")
	fmt.Scanln(&m)

	
	for n := 1; n <= m; n++ {
	
		n3 := n * n * n

	
		sum := 0
		for i := 1; sum < n3; i += 2 {
			sum += i
	
			if sum == n3 {
				fmt.Printf("%d * %d * %d = ", n, n, n)
				for j := 1; j <= i; j += 2 {
					if j > 1 {
						fmt.Print(" + ")
					}
					fmt.Print(j)
				}
				fmt.Println()
			}
		}
	}
}

func questao20() {
	var n int


	fmt.Println("Digite um número inteiro positivo:")
	fmt.Scanln(&n)


	somaDivisores := 0


	divisores := ""


	for i := 1; i < n; i++ {
		if n%i == 0 {
			somaDivisores += i
			divisores += fmt.Sprintf("%d + ", i)
		}
	}


	divisores = divisores[:len(divisores)-2]


	var mensagem string
	if somaDivisores == n {
		mensagem = "NUMERO PERFEITO"
	} else {
		mensagem = "NUMERO NAO E PERFEITO"
	}


	fmt.Printf("%d = %s = %d (%s)\n", n, divisores, somaDivisores, mensagem)
}

func questao21() {
    for {
        
        var tamanho int
        fmt.Scanln(&tamanho)
        
        
        if tamanho == 0 {
            break
        }
        
        
        sequencia := make([]float64, tamanho)
        for i := 0; i < tamanho; i++ {
            fmt.Scan(&sequencia[i])
        }
        
        
        ordenada := true
        for i := 1; i < tamanho; i++ {
            if sequencia[i-1] >= sequencia[i] {
                ordenada = false
                break
            }
        }
        
        
        if ordenada {
            fmt.Println("ORDENADA")
        } else {
            fmt.Println("DESORDENADA")
        }
    }
}
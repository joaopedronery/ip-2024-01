package main

import (
	"fmt"
	"math"
)

func main() {
	questao20();
}

func questao1() {
	var n1, n2, n3 int; 
	var media float32;
	fmt.Scanln(&n1, &n2, &n3);
	media = float32(n1 + n2 + n3) / 3;
	fmt.Printf("Media = %.2f \n", media );
	if media >= 6 {
		fmt.Println("Aprovado");
	} else {
		fmt.Println("Reprovado");
	}
}

func questao2() {
	var casosTeste int;
	var messages []string;
	fmt.Println("Informe o numero de casos de teste:")
	fmt.Scanln(&casosTeste);
	for i := 1; i <= casosTeste; i++ {
		var publicoTotal, percentPopular, percentGeral, percentArquibancada, percentCadeira int;
		fmt.Println("Digite, separados por um espaço entre cada um, o publico total, a porcentagem de publico na Popular, a porcentagem na Geral, a porcentagem na Arquibancada e a porcentagem na Cadeiras:");
		fmt.Scanln(&publicoTotal, &percentPopular, &percentGeral, &percentArquibancada, &percentCadeira);
		rendaPopular := publicoTotal * percentPopular / 100; rendaGeral := publicoTotal * percentGeral / 100 * 5; rendaArquibancada := publicoTotal * percentArquibancada / 100 * 10; rendaCadeira := publicoTotal * percentCadeira / 100 * 20;
		rendaTotal := rendaPopular + rendaArquibancada + rendaGeral + rendaCadeira;
		message := fmt.Sprintf("A RENDA DO JOGO N.%d E = %.2f", i, float32(rendaTotal));
		messages = append(messages, message);
	}
	for i := range messages {
		fmt.Println(messages[i]);
	}
}

func questao3() {
	var n1, n2, n3 int;
	fmt.Scanln(&n1);
	if n1 > 9 {
		fmt.Println("DIGITO INVALIDO");
		return;
	}
	fmt.Scanln(&n2);
	if n2 > 9 {
		fmt.Println("DIGITO INVALIDO");
		return;
	}
	fmt.Scanln(&n3);
	if n3 > 9 {
		fmt.Println("DIGITO INVALIDO");
		return;
	}
	numeroFinal := n1 * 100 + n2 * 10 + n3;
	fmt.Printf("%d, %d \n", numeroFinal, numeroFinal * numeroFinal);
}

func questao4() {
	var salMinimo, quantKw float32;
	fmt.Scanln(&salMinimo, &quantKw);
	custoPorKw := float32(7) / float32(1000) * salMinimo;
	valorTotal := custoPorKw * quantKw;
	valorComDesconto := valorTotal * 0.9;
	fmt.Printf("Custo por kW: R$ %.2f \n", custoPorKw);
	fmt.Printf("Custo do consumo: R$ %.2f \n", valorTotal);
	fmt.Printf("Custo com desconto: R$ %.2f \n", valorComDesconto);
}

func questao5() {
	var contaId int;
	var valorConta, metrosCubicos float32;
	var tipoConsumidor string;
	fmt.Scanln(&contaId, &metrosCubicos, &tipoConsumidor);
	if tipoConsumidor == "R" {
		valorConta = 5 + 0.05 * metrosCubicos;
	}
	if tipoConsumidor == "C" {
		valorConta = 500;
		if metrosCubicos > 80 {
			valorConta = valorConta + 0.25 * (metrosCubicos - 80);
		}
	}
	if tipoConsumidor == "I" {
		valorConta = 800;
		if metrosCubicos > 100 {
			valorConta = valorConta + 0.04 * (metrosCubicos - 100);
		}
	}
	fmt.Printf("CONTA = %d \n", contaId);
	fmt.Printf("VALOR DA CONTA = %.2f \n", valorConta);
}

func questao6() {
	var valoresTemperaturaF []float32;
	var numValoresLidos int;

	fmt.Println("Digite o número de valores a serem lidos:")
	fmt.Scanln(&numValoresLidos);

	fmt.Println("Digite os valores:")
	for i := 0; i < numValoresLidos; i++ {
		var tempF float32;
		fmt.Scanln(&tempF);
		valoresTemperaturaF = append(valoresTemperaturaF, tempF);
	}

	for _, tempF := range valoresTemperaturaF {
		tempC := 5 * (tempF - 32) / 9;
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS \n", tempF, tempC);
	}
}

func questao7() {
	var temperaturaF, quantidadeChuvaPolegadas float32;
	fmt.Println("Digite um valor em Fahrenheit e uma quantidade de chuva em polegadas:");
	fmt.Scan(&temperaturaF, & quantidadeChuvaPolegadas);
	temperaturaC := (5 * temperaturaF - 160) / 9;
	quantidadeChuvaMilimetros := quantidadeChuvaPolegadas * 25.4;
	fmt.Printf("O VALOR EM CELSIUS = %.2f \n", temperaturaC);
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f \n", quantidadeChuvaMilimetros);
}

func questao8() {
	var raio, altura float32;
	const pi = 3.14159;
	fmt.Printf("Digite o raio e a altura da lata, em metros \n");
	fmt.Scan(&raio, &altura);
	areaC := pi * raio * raio;
	areaL := 2 * pi * raio * altura;
	areaTotal := 2 * areaC + areaL;
	custo := areaTotal * 100;
	fmt.Printf("O VALOR DO CUSTO E = %.2f \n", custo);
}

func questao9() {
	var a, b, c float32;
	fmt.Println("Digite os coeficientes A, B e C de uma equação do segundo grau");
	fmt.Scan(&a, &b, &c);
	delta := b * b - 4 * a * c;
	fmt.Printf("O VALOR DE DELTA E = %.2f \n", delta);
}

func questao10() {
	var a, b, c, d float32;
	fmt.Println("Digite os elementos a, b, c e d que formam uma matriz quadrada bidimensional");
	fmt.Scan(&a, &b, &c, &d);
	determinante := a * d - b * c;
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f \n", determinante);
}

func questao11() {
	var numero int;
	fmt.Scanln(&numero);
	if numero % 3 == 0 && numero % 5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL");
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL");
	}
}

func questao12() {
	var numeroHoras int;
	fmt.Scanln(&numeroHoras);
	horasExtras := numeroHoras % 3;
	blocos3Horas := (numeroHoras - horasExtras) / 3;
	total := float32(blocos3Horas) * 10 + float32(horasExtras) * 5;
	fmt.Printf("O VALOR A PAGAR E = %.2f \n", total);
}

func questao13() {
	var nota float32;
	var conceito string;
	fmt.Scanln(&nota);
	if nota >= 0 && nota < 6 {
		conceito = "D";
	} else if nota >= 6 && nota < 7.5 {
		conceito = "C";
	} else if nota >= 7.5 && nota < 9 {
		conceito = "B";
	} else if nota >= 9 && nota <= 10 {
		conceito = "A";
	} else {
		conceito = "Nota invalida"
	}
	fmt.Printf("NOTA = %.1f CONCEITO = %s \n", nota, conceito);
}

func questao14() {
	var altura, aresta float64;
	fmt.Scan(&altura, &aresta);
	areaBase := (3 * aresta * aresta * math.Sqrt(3)) / 2;
	volume := 1.0 / 3.0 * areaBase * altura;
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS \n", volume);
}

func questao15() {
	var numero int;
	fmt.Scanln(&numero);
	for i := 2; i <= numero; i = i + 2 {
		fmt.Printf("%d ^ %d = %d \n", i, i, i * i);
	}
}

func questao16() {
	var salario, salarioReajustado float64;
	fmt.Scanln(&salario);
	if salario <= 300 {
		salarioReajustado = salario * 1.5;
	} else {
		salarioReajustado = salario * 1.3;
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f \n", salarioReajustado);
}

func questao17() {
	var numero, quantNumerosPrintados int;
	fmt.Scan(&numero, &quantNumerosPrintados);
	if numero % 2 == 0 {
		for i := 0; i < quantNumerosPrintados; i++ {
			fmt.Printf("%d ", numero);
			numero = numero + 2;
		}
		fmt.Println();
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR");
	}
}

func questao18() {
	var numeroInicial, razao, numeroElementos int;
	fmt.Scan(&numeroInicial, &razao, &numeroElementos);
	somatorio := numeroInicial;
	for i := 0; i < numeroElementos - 1; i++ {
		numeroInicial = numeroInicial + razao;
		somatorio = somatorio + numeroInicial;
	}
	fmt.Println(somatorio);
}

func questao19() {
	var numero int;
	var somatorio float64;
	fmt.Scanln(&numero);
	if (numero < 1) {
		fmt.Println("Numero invalido!")
		return;
	}
	for divisor := 1; divisor <= numero; divisor++ {
		somatorio = somatorio + 1 / float64(divisor);
	}
	fmt.Printf("%.6f\n", somatorio);
}

func questao20() {
	var horas, minutos, segundos int;
	fmt.Scan(&horas, &minutos, &segundos);
	totalSegundos := horas * 3600 + minutos * 60 + segundos;
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos);
}
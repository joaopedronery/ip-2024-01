package main

import "fmt"

func main() {
	questao4();
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
	var  iteracoes int;
	var numero, fatorInicialOriginal, incremento, fatorInicialClone float64;

	fmt.Scan(&numero, &fatorInicialOriginal, &iteracoes, &incremento);

	fatorInicialClone = fatorInicialOriginal;
	fmt.Println("Tabuada de soma:");
	for i := 0; i < iteracoes; i++ {
		fmt.Println()
	}
}
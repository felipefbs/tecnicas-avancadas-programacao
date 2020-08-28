package tap

import "fmt"

// Q1018 https://www.urionlinejudge.com.br/judge/pt/problems/view/1018
func Q1018() {

	var number int
	var cem, cinquenta, vinte, dez, cinco, dois, um int

	fmt.Scanf("%d", &number)

	centenaMilhar := number / 100000 * 100000
	dezenaMilhar := number % 100000 / 10000 * 10000
	unidadeMilhar := number % 10000 / 1000 * 1000

	centenas := number % 1000 / 100 * 100
	dezenas := number % 100 / 10 * 10
	unidade := number % 10

	for centenaMilhar > 0 {
		centenaMilhar -= 100
		cem++
	}
	for dezenaMilhar > 0 {
		dezenaMilhar -= 100
		cem++
	}
	for unidadeMilhar > 0 {
		unidadeMilhar -= 100
		cem++
	}
	for centenas > 0 {
		centenas -= 100
		cem++
	}
	for dezenas >= 50 {
		dezenas -= 50
		cinquenta++
	}
	for dezenas >= 20 {
		dezenas -= 20
		vinte++
	}
	for dezenas >= 10 {
		dezenas -= 10
		dez++
	}
	for unidade >= 5 {
		unidade -= 5
		cinco++
	}
	for unidade >= 2 {
		unidade -= 2
		dois++
	}
	for unidade >= 1 {
		unidade--
		um++
	}

	fmt.Printf("%d\n%d nota(s) de R$ 100,00\n%d nota(s) de R$ 50,00\n%d nota(s) de R$ 20,00\n%d nota(s) de R$ 10,00\n%d nota(s) de R$ 5,00\n%d nota(s) de R$ 2,00\n%d nota(s) de R$ 1,00\n", number, cem, cinquenta, vinte, dez, cinco, dois, um)
}

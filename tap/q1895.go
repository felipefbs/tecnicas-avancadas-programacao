package tap

import (
	"fmt"
)

// Q1895 https://www.urionlinejudge.com.br/judge/pt/problems/view/1895
func Q1895() {
	var pontosA, pontosB, limite, nCartas, i, cartaInicial, valor int8
	var arrayCartas []int8

	fmt.Scanf("%d %d %d", &nCartas, &cartaInicial, &limite)

	cartaMesa := cartaInicial

	for i = 0; i < nCartas-1; i++ {
		fmt.Scanf("%d", &valor)
		arrayCartas = append(arrayCartas, valor)
	}

	for i, v := range arrayCartas {
		dif := func(n int8) int8 {
			y := n >> 7
			return (n ^ y) - y
		}(v - cartaMesa)

		if dif <= limite {
			if i%2 == 0 {
				pontosA += dif
				cartaMesa = v
			}
			if i%2 != 0 {
				pontosB += dif
				cartaMesa = v
			}
		}
	}
	fmt.Printf("%d %d\n", pontosA, pontosB)
}

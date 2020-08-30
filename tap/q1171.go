package tap

import (
	"fmt"
	"sort"
)

// Q1171 https://www.urionlinejudge.com.br/judge/pt/problems/view/1171
func Q1171() {
	var entrada int
	var arrayEntrada []int
	var indices []int

	mapInfo := make(map[int]int)

	fmt.Scanf("%d", &entrada)

	for i := 0; i < entrada; i++ {
		var valor int
		fmt.Scanf("%d", &valor)
		arrayEntrada = append(arrayEntrada, valor)
	}

	for _, v := range arrayEntrada {

		_, existe := mapInfo[v]

		if !existe {
			mapInfo[v] = 1
			indices = append(indices, v)
		} else {
			mapInfo[v] = mapInfo[v] + 1
		}
	}

	sort.Ints(indices)

	for _, i := range indices {
		fmt.Printf("%d aparece %d vez(es)\n", i, mapInfo[i])
	}
}

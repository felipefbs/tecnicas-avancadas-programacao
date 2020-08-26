package tap

import "fmt"

// Q1018 https://www.urionlinejudge.com.br/judge/pt/problems/view/1157
func Q1157() {
	var number int

	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

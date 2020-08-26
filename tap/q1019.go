package tap

import "fmt"

// Q1019 https://www.urionlinejudge.com.br/judge/pt/problems/view/1019
func Q1019() {
	var number, hours, minutes, seconds int

	fmt.Scanf("%d", &number)

	seconds = number % 60

	minutes = int(number / 60)
	if minutes >= 60 {
		hours = int(minutes / 60)
		minutes = minutes % 60
	}

	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)

}

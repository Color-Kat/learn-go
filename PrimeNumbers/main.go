package main

import (
	"fmt"
	"math"
)

func addPrint(num int) {
	fmt.Printf("%v ", num)
}

func main() {
	max := int(math.Pow(2, 32) + 1)
	for n := 0; n < max; n++ {
		if n == 2 {
			addPrint(2)
			continue
		}

		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		// if isPrime {
		// 	addPrint(n)
		// }

		if isPrime && n == int(math.Pow(2, 32)-1) {
			fmt.Print("Yes - ")
			fmt.Print(math.Pow(2, 32) - 1)
		}

	}
}

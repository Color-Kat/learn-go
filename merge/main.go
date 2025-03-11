package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ToInt = strconv.Atoi

func Prompt() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func main() {

	//n, _ := ToInt(Prompt())
	//m, _ := ToInt(Prompt())

	//for i := 0; i < n; i++ {
	//	a[i], _ = ToInt(Prompt())
	//	a[i], _ = ToInt(Prompt())
	//}
	//
	//for i := 0; i < m; i++ {
	//	b[i], _ = ToInt(Prompt())
	//}

	a := []int{1, 2, 5, 6, 84}
	b := []int{2, 3, 6, 7, 40, 85}
	n := len(a)
	m := len(b)
	c := make([]int, n+m)

	l, r := 0, 0
	for l < n || r < m {
		if l < n && r < m {
			if a[l] < b[r] {
				c[l+r] = a[l]
				l++
			} else {
				c[l+r] = b[r]
				r++
			}
		} else {
			for l < n {
				c[l+r] = a[l]
				l++
			}

			for r < m {
				c[l+r] = b[r]
				r++
			}
		}
	}

	fmt.Println(c)
}

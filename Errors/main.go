package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (divideError divideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero", divideError.dividend)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, divideError{dividend: a}
	}

	return a / b, nil
}

func main() {
	for i := 0; i < 10; i++ {
		for j := -5; j < 5; j++ {
			_, err := divide(float64(i), float64(j))

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

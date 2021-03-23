package calculator

import "fmt"

var message = "LOG"

var Version = "3.0"

func internalSum(number int) int {
	return number + 1
}

func Sum(a int, b int) int {
	fmt.Println("Computing the sum...")
	return a + b
}

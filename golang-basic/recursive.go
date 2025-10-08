package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}
func main() {
	fmt.Println(factorial(10))
	fmt.Println(factorialRecursive(10))
}

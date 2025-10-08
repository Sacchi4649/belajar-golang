package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	preparedNumbers := []int{10, 10, 30}
	fmt.Println(sumAll(10, 20, 30, 40, 50))
	fmt.Println(sumAll(preparedNumbers...))
}

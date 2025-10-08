package main

import "fmt"

func Ups() any {
	return "Ups"
}

func main() {
	any := Ups()
	fmt.Println(any)
}

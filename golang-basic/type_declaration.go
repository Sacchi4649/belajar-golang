package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NIK NoKTP = "1234567890"
	var MarriedStatus Married = true

	fmt.Println(NIK, MarriedStatus)

	MarriedStatus = false
	fmt.Println(NIK, MarriedStatus)
}

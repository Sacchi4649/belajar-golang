package main

import "fmt"

// func getFullName(firstName string, lastName string) (string, string) {
// 	return firstName, lastName
// }

func getFullName() (string, string) {
	return "Sany", "Adika"
}

func main() {
	_, lastName := getFullName()
	fmt.Println(lastName)
}

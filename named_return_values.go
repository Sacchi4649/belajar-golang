package main

import "fmt"

func getFullBiodata() (firstName, lastName string, age int) {

	firstName = "Sany"
	lastName = "Adika"
	age = 20

	return firstName, lastName, age
}

func main() {
	firstName, lastName, age := getFullBiodata()
	fmt.Println(firstName, lastName, age)
}

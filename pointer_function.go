package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	// var address *Address = &Address{}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}

//pointer function itu function yang mengembalikan pointer

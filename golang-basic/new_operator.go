package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	alamat2 := alamat1
	alamat1.City = "Jakarta"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

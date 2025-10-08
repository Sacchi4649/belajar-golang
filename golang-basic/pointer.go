package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//var address2 *Address = &address1
	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 := &address1
	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)
}

//pointer itu referensi dari value lain
//kalau tidak pakai pointer, maka akan membuat copy dari value tersebut

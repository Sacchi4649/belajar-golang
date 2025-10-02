package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Sany"
	names[1] = "Dika"
	names[2] = "Ima"

	fmt.Println(names[0], names[1], names[2])

	names[0] = "Adi"
	fmt.Println(names[0], names[1], names[2])

	values := [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	values[0] = 100
	fmt.Println(values)

	fmt.Println("Panjang array names", len(names))
	fmt.Println(len(values))

	dynamicArray := [...]int{
		1,
		2,
		3,
	}
	fmt.Println(dynamicArray)
	fmt.Println(len(dynamicArray))

}

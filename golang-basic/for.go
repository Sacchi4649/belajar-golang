package main

import "fmt"

func main() {
	names := []string{"Sany", "Dika", "Ima"}

	// for i, name := range names {
	// 	fmt.Println("index", i, "name", name)
	// }

	for _, name := range names {
		fmt.Println("name", name)
	} // kalau tidak butuh value index, bisa gunakan _

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
}

//counter := 1 itu init statement
//counter++ itu post statement (selalu dilakukan di akhir perulangan)

// hitung:=1
// for hitung <=10{
// 	fmt.Println("Perulangan ke", hitung)
// 	hitung++
// }

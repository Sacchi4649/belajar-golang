package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
	// 	"name":    "Sany",
	// 	"address": "Jakarta",
	// }

	person := map[string]string{
		"name":    "Sany",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Sany"
	book["ups"] = "Salah"

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["ups"])
}

package main

import "fmt"

func main() {
	umur := 20
	name := "Sany22"
	if umur > 18 {
		fmt.Println("Anda sudah mencukupi umur")
	} else {
		fmt.Println("Anda belum mencukupi umur")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}

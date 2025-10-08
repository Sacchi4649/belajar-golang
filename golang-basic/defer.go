package main

import "fmt"

func main() {
	defer fmt.Println("Selesai")
	fmt.Println("Hello")
}

//defer itu dipanggil diakhir

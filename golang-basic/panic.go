package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Terjadi panic")
	}
}
func main() {
	runApp(true)
}

//panic itu untuk menghentikan program tapi defer tetap akan dijalankan
//recover itu untuk menangkap panic, diletakkan pada defer function

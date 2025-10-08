package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	sany := Man{Name: "Sany"}
	sany.Married()
	fmt.Println(sany.Name)
}

//untuk method, direkomendasikan untuk menggunakan pointer jika ingin mengubah value dari struct

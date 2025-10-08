package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	contoh := goodBye
	fmt.Println(goodBye("Ikhsan"), contoh("Sany"))
}

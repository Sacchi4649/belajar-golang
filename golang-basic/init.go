package main

import (
	"fmt"
	"hello-world/database"
	_ "hello-world/internal" //kalau cuma mau init, tapi ga mau  pakai function di package tersebut. pakai _ sebagai blank identifier
)

func main() {
	fmt.Println(database.GetConnection())
}

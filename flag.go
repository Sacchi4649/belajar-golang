package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Database username")
	password := flag.String("password", "root", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")
	database := flag.String("database", "mydatabase", "Database name")
	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
	fmt.Println("Database:", *database)
}

//contoh penggunaan flag
//go run flag.go -username=sany -password="secret_password"

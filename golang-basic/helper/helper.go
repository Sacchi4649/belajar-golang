package helper

import "fmt"

var Application = "Hello World"
var Version = "1.0.0"

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}

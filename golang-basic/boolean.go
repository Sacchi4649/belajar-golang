package main
import "fmt"

func main(){
	fmt.Println("benar", true)
	fmt.Println("salah", false)
	fmt.Println("benar sama dengan salah", true == false)
	fmt.Println("benar sama dengan benar", true == true)
	fmt.Println("salah sama dengan salah", false == false)
	fmt.Println("salah sama dengan benar", false == true)
}
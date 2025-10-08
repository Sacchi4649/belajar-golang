package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	return nilai / pembagi, nil
}

func main() {
	result, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err.Error())
	}

}

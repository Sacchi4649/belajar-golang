package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID tidak boleh kosong"}
	}
	if id != "sany" {
		return &notFoundError{Message: "Data tidak ditemukan"}
	}
	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil {
		// 	if validationError, ok := err.(*validationError); ok {
		// 		fmt.Println("Validation Error", validationError.Error())
		// 	} else if notFoundError, ok := err.(*notFoundError); ok {
		// 		fmt.Println("Not Found Error", notFoundError.Error())
		// 	} else {
		// 		fmt.Println("Error", err.Error())
		// 	}
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error", finalError.Error())
		default:
			fmt.Println("Error", finalError.Error())
		}
	} else {
		fmt.Println("Data berhasil disimpan")
	}

}

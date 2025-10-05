package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func (user User) GetName() string {
	return user.Name
}

func (user User) GetEmail() string {
	return user.Email
}

func (user User) GetAge() int {
	return user.Age
}

func (user User) GetID() int {
	return user.ID
}

func (user User) GetAll() (int, string, string, int) {
	return user.ID, user.Name, user.Email, user.Age
}

func main() {
	user1 := User{
		ID:    1,
		Name:  "Sany",
		Email: "sany@gmail.com",
		Age:   20,
	}
	Dika := User{
		ID:    2,
		Name:  "Dika",
		Email: "dika@gmail.com",
		Age:   21,
	}
	Mama := User{3, "mama", "mama@gmail.com", 22}
	fmt.Println(user1.GetName())

	fmt.Println(user1.Age)
	fmt.Println(Dika)
	fmt.Println(Mama)
}

//struct itu adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//struct adalah template data

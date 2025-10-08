package main

import "fmt"

func main() {
	names := [...]string{
		"Sany",
		"Dika",
		"Ima",
		"Ibnul",
		"Khairul",
		"Aziz",
	}
	// var slice2 []string = names[:]

	namesSlice := names[2:] // slice dari index 2 sampai akhir
	nameSlice2 := names[:4] // slice dari index 0 sampai sebelum index 4
	// kalau names[:6] slice dari index 0 sampai 6
	// kalau names[:] slice dari index 0 sampai akhir
	// kalau names[2:6] slice dari index 2 sampai sebelum index 6
	// kalau names[2:] slice dari index 2 sampai akhir
	// kalau names[:6] slice dari index 0 sampai sebelum index 6
	// kalau names[:] slice dari index 0 sampai akhir
	// kalau names[2:6] slice dari index 2 sampai 6
	fmt.Println(namesSlice)
	fmt.Println(nameSlice2)
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice := days[2:5]
	fmt.Println(daysSlice)

	daysSlice[0] = "Rabu Baru"
	fmt.Println(daysSlice)
	fmt.Println(days)

	daysSlice = append(daysSlice, "Ahad")
	fmt.Println("Setelah append ", daysSlice)
	fmt.Println(days)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

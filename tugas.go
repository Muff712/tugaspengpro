package main

import "fmt"

func main() {
	var a string

	fmt.Println("pilih pengguna : ")
	fmt.Println(" Admin ")
	fmt.Println(" User ")

	fmt.Scan(&a)

	switch a {
	case "Admin":
		menuadmin()
	case "User":
		menuuser()
	case "keluar":
		fmt.Println("anda keluar")
	}
}

func menuadmin() {
	var a int
	fmt.Println("---MENU---")
	fmt.Println("1. Input data ")
	fmt.Println("2. hapus data ")
	fmt.Println("2. hapus data ")

	fmt.Scan(&a)

	switch a {
	case 1:
		Inputdata()
	case 2:
		hapusdata()
	}
}

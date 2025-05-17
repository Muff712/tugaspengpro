package main

import "fmt"

type tim struct {
	nama                                                                                           string
	jumlahpertandingan, peringkat, jumlahmenang, jumlahseri, jumlahkalah, jumlahanggota, jumlahtim int
	macthrate                                                                                      float32
}

type pemain struct {
	nama                                                                                string
	jumlahpertandingan, peringkat, jumlahmenang, jumlahseri, jumlahkalah, jumlahanggota int
	posisi                                                                              string
}
type pertandingan struct {
	tim1, tim2 tim.nama
}

type tm [100 - 1]tim
type pm [100 - 1]pemain
type pt [100 - 1]pertandingan

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
	fmt.Println("3. keluar ")

	fmt.Scan(&a)

	switch a {
	case 1:
		inputdata()
	case 2:
		hapusdata()
	}
}

func inputdata() {
	var a int

	fmt.Println("---MENU---")
	fmt.Println("1. Input data tim ")
	fmt.Println("2. input data pemain")
	fmt.Println("3. input data pertandingan")
	fmt.Println("4. keluar ")
	fmt.Scan(&a)

	switch a {
	case 1:
		inputtim()
	case 2:
		inputpemain()
	case 3:
		inputpertandingan()
	}
}
func hapusdata() {
	var a int

	fmt.Println("---MENU---")
	fmt.Println("1. Hapus data tim ")
	fmt.Println("2. hapus data pemain")
	fmt.Println("3. hapus data pertandingan")
	fmt.Println("4. keluar ")
	fmt.Scan(&a)

	switch a {
	case 1:
		hapustim()
	case 2:
		hapuspemain()
	case 3:
		hapuspertandingan()
	}
}

func inputtim(a *tm, b *int) {
	fmt.Print("masukkan nama tim : ")
	fmt.Scan(&b)
	for i := 0; i < *b; i++ {
		fmt.Print("masukkan nama tim : ")
		fmt.Scan(&a[i].nama)
		fmt.Print("masukkan jumlah anggota : ")
		fmt.Scan(&a[i].jumlahanggota)
	}
}

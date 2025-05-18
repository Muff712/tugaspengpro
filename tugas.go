package main

import "fmt"

type tim struct {
	nama                                                                                           string
	jumlahpertandingan, peringkat, jumlahmenang, jumlahseri, jumlahkalah, jumlahanggota, jumlahtim int
	macthrate                                                                                      float32
}

type pemain struct {
	nama, namatim                                                                       string
	jumlahpertandingan, peringkat, jumlahmenang, jumlahseri, jumlahkalah, jumlahanggota int
	posisi                                                                              string
}
type pertandingan struct {
	tim1, tim2 string
}

var totaltim int

type tm [100 - 1]tim
type pm [100 - 1]pemain
type pt [100 - 1]pertandingan

func menu(a string) {

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
	fmt.Println("3. tampilkan data ")
	fmt.Println("3. keluar ")

	fmt.Scan(&a)

	switch a {
	case 1:
		inputdata()
	case 2:
		hapusdata()
	case 3:
		tampildata()
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
	fmt.Println("4. keluar ")
	fmt.Scan(&a)

	switch a {
	case 1:
		hapustim()
	case 2:
		hapuspemain()
	}
}

func inputtim(a *tm, b *int) {
	var v string
	fmt.Print("masukkan jumlah tim : ")
	fmt.Scan(&b)
	for i := 0; i < *b; i++ {
		fmt.Print("masukkan nama tim : ")
		fmt.Scan(&a[i].nama)
		fmt.Print(" riwayat pertandingan tim : ")
		fmt.Scan(&v)
		if v == "ya" {
			fmt.Print("masukkan jumlah pertandingan : ")
			fmt.Scan(&a[i].jumlahpertandingan)
			fmt.Print("masukkan jumlah menang : ")
			fmt.Scan(&a[i].jumlahmenang)
			fmt.Print("masukkan jumlah seri : ")
			fmt.Scan(&a[i].jumlahseri)
			fmt.Print("masukkan jumlah kalah : ")
			fmt.Scan(&a[i].jumlahkalah)
		} else {
			a[i].jumlahpertandingan = 0
			a[i].jumlahmenang = 0
			a[i].jumlahseri = 0
			a[i].jumlahkalah = 0
		}
	}
	totaltim = totaltim + *b
}

func inputpemain(a *pm, b *int, c *string) {
	var e string
	fmt.Print("masukkan nama tim : ")
	fmt.Scan(&c)
	fmt.Print("masukkan jumlah pemain : ")
	fmt.Scan(&b)
	for i := 0; i < *b; i++ {
		fmt.Print("masukkan nama pemain : ")
		fmt.Scan(&a[i].nama)
		fmt.Print("masukkan posisi pemain : ")
		fmt.Scan(&a[i].posisi)
		fmt.Print("riwayat pertandingan pemain : ")
		fmt.Scan(&e)
		if e == "ya" {
			fmt.Print("masukkan jumlah pertandingan : ")
			fmt.Scan(&a[i].jumlahpertandingan)
			fmt.Print("masukkan jumlah menang : ")
			fmt.Scan(&a[i].jumlahmenang)
			fmt.Print("masukkan jumlah seri : ")
			fmt.Scan(&a[i].jumlahseri)
			fmt.Print("masukkan jumlah kalah : ")
			fmt.Scan(&a[i].jumlahkalah)
		} else {
			a[i].jumlahpertandingan = 0
			a[i].jumlahmenang = 0
			a[i].jumlahseri = 0
			a[i].jumlahkalah = 0
		}
		a[i].namatim = *c
	}
}

func hapustim(a *tm, b *int) {
	var c string
	fmt.Print("masukkan nama tim yang ingin dihapus : ")
	fmt.Scan(&c)
	for i := 0; i < *b; i++ {
		if a[i].nama == c {
			a[i] = ""
			fmt.Println("data tim berhasil dihapus")
		} else {
			fmt.Println("data tim tidak ditemukan")
		}
	}
}

func hapuspemain(a *pm, b *int) {
	var c string
	fmt.Print("masukkan nama pemain yang ingin dihapus : ")
	fmt.Scan(&c)
	for i := 0; i < *b; i++ {
		if a[i].nama == c {
			a[i] = pemain{}
			fmt.Println("data pemain berhasil dihapus")
		} else {
			fmt.Println("data pemain tidak ditemukan")
		}
	}
}

func inputpertandingan(a *pt, b *int) {
	var d int
	fmt.Print("masukkan jumlah pertandingan : ")
	fmt.Scan(&b)
	d = 1
	for i := 0; i < *b; i++ {
		fmt.Print("masukkan pertandingan ke-", d, " : ")
		d++
		fmt.Print("masukkan nama tim 1 : ")
		fmt.Scan(&a[i].tim1)
		fmt.Print("masukkan nama tim 2 : ")
		fmt.Scan(&a[i].tim2)
	}
}

func menuuser() {
	var a int
	fmt.Println("---MENU---")
	fmt.Println("1. Tampilkan data tim ")
	fmt.Println("2. tampilkan data pemain")
	fmt.Println("3. tampilkan data pertandingan")
	fmt.Println("4. keluar ")
	fmt.Scan(&a)

	switch a {
	case 1:
		tampiltim()
	case 2:
		tampilpemain()
	case 3:
		tampilpertandingan()
	}
}

func tampiltim(a *tm, b *int) {
	fmt.Println("data tim : ")
	for i := 0; i < *b; i++ {
		fmt.Println(" %s ", a[i].nama)
	}
}

func tampilpemain(a *pm, b *int) {
	fmt.Println("data pemain : ")
	for i := 0; i < *b; i++ {
		fmt.Println(" %s ", a[i].nama)
	}
}

//JadwalOtomatis
func jadwalOtomatis(timData *tm, jumlahTim int, jadwal *pt, jumlahJadwal *int) {
	for i := 0; i < jumlahTim; i++ {
		for j := i + 1; j < jumlahTim; j++ {
			jadwal[*jumlahJadwal].tim1 = timData[i].nama
			jadwal[*jumlahJadwal].tim2 = timData[j].nama
			(*jumlahJadwal)++
		}
	}
	fmt.Println("Jadwal pertandingan otomatis telah dibuat.")
}

//Statistik
func tampilStatistikTim(data *tm, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Tim: %s | Main: %d | Menang: %d | Seri: %d | Kalah: %d\n",
			data[i].nama, data[i].jumlahpertandingan, data[i].jumlahmenang, data[i].jumlahseri, data[i].jumlahkalah)
	}
}

func tampilStatistikPemain(data *pm, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Pemain: %s | Tim: %s | Posisi: %s | Main: %d\n",
			data[i].nama, data[i].namatim, data[i].posisi, data[i].jumlahpertandingan)
	}
}

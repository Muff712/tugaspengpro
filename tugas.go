package main

import (
	"fmt"
	"strings"
)

const (
	MaxTeams   = 100
	MaxPlayers = 100
	MaxMatches = 100
)

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
	hari       string
	jam        int
}

var totaltim int

type tm [100 - 1]tim
type pm [100 - 1]pemain
type pt [100 - 1]pertandingan

type MenuOption struct {
	ID          int
	Description string
}

func menu() error {
	var choice string
	options := []string{"Admin", "User", "Keluar"}

	fmt.Println("Pilih pengguna:")
	for _, opt := range options {
		fmt.Printf(" %s\n", opt)
	}

	if _, err := fmt.Scan(&choice); err != nil {
		return fmt.Errorf("invalid input: %v", err)
	}

	switch strings.ToLower(choice) {
	case "admin":
		return menuadmin()
	case "user":
		return menuuser()
	case "keluar":
		fmt.Println("Anda keluar")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
}

func menuadmin() error {
	var a int
	fmt.Println("---MENU---")
	fmt.Println("1. Input data ")
	fmt.Println("2. hapus data ")
	fmt.Println("3. tampilkan data ")
	fmt.Println("4. keluar ")

	if _, err := fmt.Scan(&a); err != nil {
		return fmt.Errorf("invalid input: %v", err)
	}

	switch a {
	case 1:
		inputdata()
	case 2:
		hapusdata()
	case 3:
		menuuser()
	}
	return nil
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
		teams := &TeamList{teams: make([]tim, MaxTeams)}
		if err := inputTim(teams); err != nil {
			fmt.Println("Error:", err)
		}
	case 2:
		var players pm
		var playerCount int
		var teamName string
		inputpemain(&players, &playerCount, &teamName)
	case 3:
		var matches pt
		var matchCount int
		inputpertandingan(&matches, &matchCount)
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
		var teams tm
		hapustim(&teams, &totaltim)
	case 2:
		var players pm
		var playerCount int
		hapuspemain(&players, &playerCount)
	}
}

type TeamList struct {
	teams []tim
}

func inputTim(teams *TeamList) error {
	var count int
	fmt.Print("masukkan jumlah tim : ")
	if _, err := fmt.Scan(&count); err != nil {
		return fmt.Errorf("invalid input: %v", err)
	}
	if count <= 0 || count > MaxTeams {
		return fmt.Errorf("jumlah tim harus antara 1 dan %d", MaxTeams)
	}

	for i := 0; i < count; i++ {
		var v string
		fmt.Print("masukkan nama tim : ")
		fmt.Scan(&teams.teams[i].nama)
		fmt.Print(" riwayat pertandingan tim : ")
		fmt.Scan(&v)
		if v == "ya" {
			fmt.Print("masukkan jumlah pertandingan : ")
			fmt.Scan(&teams.teams[i].jumlahpertandingan)
			fmt.Print("masukkan jumlah menang : ")
			fmt.Scan(&teams.teams[i].jumlahmenang)
			fmt.Print("masukkan jumlah seri : ")
			fmt.Scan(&teams.teams[i].jumlahseri)
			fmt.Print("masukkan jumlah kalah : ")
			fmt.Scan(&teams.teams[i].jumlahkalah)
		} else {
			teams.teams[i].jumlahpertandingan = 0
			teams.teams[i].jumlahmenang = 0
			teams.teams[i].jumlahseri = 0
			teams.teams[i].jumlahkalah = 0
		}
	}
	totaltim = totaltim + count
	return nil
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
			a[i] = tim{}
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
	i := 0
	for i < *b {
		var konflict bool
		fmt.Print("masukkan pertandingan ke-", d, " : ")
		fmt.Print("masukkan nama tim 1 : ")
		fmt.Scan(&a[i].tim1)
		fmt.Print("masukkan nama tim 2 : ")
		fmt.Scan(&a[i].tim2)

		fmt.Print("masukkan hari (senin/selasa/rabu/kamis/jumat/sabtu/minggu): ")
		fmt.Scan(&a[i].hari)

		fmt.Print("masukkan jam (format 24 jam, 0-23): ")
		fmt.Scan(&a[i].jam)

		if i > 0 {
			for j := 0; j < i && !konflict; j++ {
				if (a[i].hari == a[j].hari && a[i].jam == a[j].jam) &&
					(a[i].tim1 == a[j].tim1 || a[i].tim1 == a[j].tim2 ||
						a[i].tim2 == a[j].tim1 || a[i].tim2 == a[j].tim2) {
					fmt.Println("PERINGATAN: Jadwal bertabrakan!")
					fmt.Println("Tim yang bertanding sudah ada di jadwal yang sama")
					konflict = true
				}
			}
		}

		if konflict != true {
			i++
			d++
		}
	}
}

func menuuser() error {
	var a int
	fmt.Println("---MENU---")
	fmt.Println("1. Tampilkan data tim ")
	fmt.Println("2. tampilkan data pemain")
	fmt.Println("3. tampilkan data pertandingan")
	fmt.Println("4. keluar ")
	fmt.Scan(&a)

	switch a {
	case 1:
		var teams tm
		tampiltim(&teams, &totaltim)
	case 2:
		var players pm
		var playerCount int
		tampilpemain(&players, &playerCount)
	case 3:
		var matches pt
		var matchCount int
		tampilpertandingan(&matches, &matchCount)
	}
	return nil
}

func tampiltim(a *tm, b *int) {
	fmt.Println("DATA DAN STATISTIK TIM:")
	fmt.Println("----------------------------------------")
	for i := 0; i < *b; i++ {
		if a[i].nama != "" {
			fmt.Printf("Nama Tim: %s\n", a[i].nama)
			fmt.Printf("Jumlah Pertandingan: %d\n", a[i].jumlahpertandingan)
			fmt.Printf("Menang: %d | Seri: %d | Kalah: %d\n",
				a[i].jumlahmenang, a[i].jumlahseri, a[i].jumlahkalah)
			if a[i].jumlahpertandingan > 0 {
				winRate := float32(a[i].jumlahmenang) / float32(a[i].jumlahpertandingan) * 100
				fmt.Printf("Win Rate: %.1f%%\n", winRate)
			}
			fmt.Println("----------------------------------------")
		}
	}
}

func tampilpemain(a *pm, b *int) {
	fmt.Println("DATA DAN STATISTIK PEMAIN:")
	fmt.Println("----------------------------------------")
	for i := 0; i < *b; i++ {
		if a[i].nama != "" {
			fmt.Printf("Nama Pemain: %s\n", a[i].nama)
			fmt.Printf("Tim: %s\n", a[i].namatim)
			fmt.Printf("Posisi: %s\n", a[i].posisi)
			fmt.Printf("Jumlah Pertandingan: %d\n", a[i].jumlahpertandingan)
			fmt.Printf("Menang: %d | Seri: %d | Kalah: %d\n",
				a[i].jumlahmenang, a[i].jumlahseri, a[i].jumlahkalah)
			if a[i].jumlahpertandingan > 0 {
				winRate := float32(a[i].jumlahmenang) / float32(a[i].jumlahpertandingan) * 100
				fmt.Printf("Win Rate: %.1f%%\n", winRate)
			}
			fmt.Println("----------------------------------------")
		}
	}
}

func tampilpertandingan(a *pt, b *int) {
	fmt.Println("DATA PERTANDINGAN:")
	fmt.Println("----------------------------------------")
	for i := 0; i < *b; i++ {
		fmt.Printf("Pertandingan %d:\n", i+1)
		fmt.Printf("Tim: %s vs %s\n", a[i].tim1, a[i].tim2)
		fmt.Printf("Jadwal: Hari %s, Jam %d:00\n", a[i].hari, a[i].jam)
		fmt.Println("----------------------------------------")
	}
}

func main() {
	for {
		if err := menu(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

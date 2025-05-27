package main

import "fmt"


type Tim struct {
	Nama                                                                                string
	JumlahPertandingan, Peringkat, JumlahMenang, JumlahSeri, JumlahKalah, JumlahAnggota int
	RasioMenang                                                                         float32
}

type TabTim [8 - 1]Tim


type Pemain struct {
	Nama               string
	Posisi             string
}

type TabPemain [10 - 1]Pemain


type Pertandingan struct {
	Tim1, Tim2  string
	Hari        string
	HasilTim1   int
	HasilTim2   int
}

type TabPertandingan [9 - 1]Pertandingan

type Minggu [15 - 1]int


func Menu() {
	var pilihan string
	fmt.Println("Pilih pengguna:")
	fmt.Println(" Admin")
	fmt.Println(" User")
	fmt.Println(" Keluar")
	fmt.Scan(&pilihan)

	if pilihan == "admin" || pilihan == "Admin" {
		MenuAdmin()
	} else if pilihan == "user" || pilihan == "User" {
		MenuPengguna()
	} else if pilihan == "keluar" || pilihan == "Keluar" {
		fmt.Println("Anda keluar")
		return
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func MenuAdmin() {
	var pilihan int
	fmt.Println("---MENU ADMIN---")
	fmt.Println("1. Input Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Tampilkan Data")
	fmt.Println("4. Keluar")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		InputData()
	case 2:
		HapusData()
	case 3:
		MenuPengguna()
	case 4:
		return
	case 5:
		Menu()
	}
	fmt.Println("Pilihan tidak valid")
	MenuAdmin()
}

func InputData() {
	var pilihan int
	fmt.Println("---MENU INPUT---")
	fmt.Println("1. Input Data Tim")
	fmt.Println("2. Input jadwal")
	fmt.Println("3. Input skor pertandingan")
	fmt.Println("4. Kembali ke Menu Admin")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		InputTim()
	case 2:
		InputPertandingan()
	case 3:
		InputSkor()
	case 4:
		MenuAdmin()
	}
	fmt.Println("Pilihan tidak valid")
	InputData()
}

func InputTim(a *TabTim) {
	var b TabPemain
	fmt.Println("=== INPUT DATA TIM ===")
	for i := 0; i < 7; i++ {
		fmt.Printf("\nData Tim ke-%d:\n")
		fmt.Println("Masukkan nama tim: ")
		fmt.Scan(&a[i].Nama)
		fmt.Print("Masukkan jumlah anggota: ")
		fmt.Scan(&a[i].JumlahAnggota)
		InputPemain(&b, a[i].JumlahAnggota)
		}
	}

func InputPemain(a *TabPemain, n int) {
	for i := 0;i < n; i++ {
		fmt.Printf("\nData Pemain ke-%d:\n", i+1)
		fmt.Println("Masukkan nama pemain: ")
		fmt.Scan(&a[i].Nama)
		fmt.Print("Masukkan posisi: ")
		fmt.Scan(&a[i].Posisi)
	}
}

func InputPertandingan(a *TabPertandingan) {
	var d, e TabTim
	fmt.Println("=== INPUT JADWAL PERTANDINGAN ===")
	for i := 0; i < 8; i++ {
		fmt.Printf("\nData Pertandingan ke-%d:\n", i+1)
		fmt.Print("Masukkan nama tim 1: ")
		fmt.Scan(&a[i].Tim1)	
		fmt.Print("Masukkan nama tim 2: ")
		fmt.Scan(&a[i].Tim2)
		b := Validtim(d, &a[i].Tim1)
		c := Validtim(e, &a[i].Tim2)
		if !b && !c {
			fmt.Println("Tim tidak valid, silakan masukkan nama tim yang benar.")
			i-- 
		}
		fmt.Print("Masukkan hari pertandingan: ")
		fmt.Scan(&a[i].Hari)
	}
}

func Validtim(a TabTim, nama *string) bool {
b := false
	for i := 0; i < 7; i++ {
		if a[i].Nama == *nama {
		return true
		}
	}
	return (b)
}

func InputSkor(a *TabPertandingan) {
	fmt.Println("=== INPUT SKOR PERTANDINGAN ===")
	for i := 0; i < 8; i++ {
		fmt.Printf("\nData Pertandingan ke-%d:\n", i+1)
		fmt.Printf("Tim 1: %s\n", a[i].Tim1)
		fmt.Printf("Tim 2: %s\n", a[i].Tim2)
		fmt.Print("Masukkan skor tim 1: ")
		fmt.Scan(&a[i].HasilTim1)
		fmt.Print("Masukkan skor tim 2: ")
		fmt.Scan(&a[i].HasilTim2)
	}
}

func CariIdxTim(a *TabTim, n int, nama string) int {
    for i := 0; i < n; i++ {
        if a[i].Nama == nama {
            return i
        }
    }
    return -1
}

func HapusData(a *TabTim, b *[7]TabPemain, n *int) {
    var nama string
    fmt.Println("=== HAPUS TIM ===")
    fmt.Print("Masukkan nama tim yang ingin dihapus: ")
    fmt.Scan(&nama)

    idx := CariIdxTim(a, *n, nama)
    if idx == -1 {
        fmt.Println("Tim tidak ditemukan.")
        return
    }

    for i := idx; i < *n-1; i++ {
        a[i] = a[i+1]
        b[i] = b[i+1]
    }
    *n--
    fmt.Println("Tim dan data pemain berhasil dihapus.")
}









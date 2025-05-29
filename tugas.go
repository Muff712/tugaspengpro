package main

import "fmt"

type Tim struct {
	Nama                                          string
	JumlahPertandingan, JumlahMenang, JumlahKalah int
	RasioMenang                                   float32
}

type TabTim [6]Tim

type Pemain struct {
	Nama   string
	Posisi string
}

type TabPemain [7]Pemain

type Pertandingan struct {
	Tim1, Tim2 string
	Hari       string
	HasilTim1  string
	HasilTim2  string
}

type TabPertandingan [9]Pertandingan

func main() {
	Menu()
}

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
		Menu()
	}
}

func MenuAdmin() {
	var pilihan int
	var tim TabTim
	var pemain [7]TabPemain
	var n int = 7
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
		HapusData(&tim, &pemain, &n)
	case 3:
		MenuPengguna()
	case 4:
		return
	case 5:
		Menu()
	default:
		fmt.Println("Pilihan tidak valid")
		MenuAdmin()
	}
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
		var tim TabTim
		InputTim(&tim)
	case 2:
		var pertandingan TabPertandingan
		InputPertandingan(&pertandingan)
	case 3:
		var pertandingan TabPertandingan
		InputHasil(pertandingan)
	case 4:
		MenuAdmin()
	default:
		fmt.Println("Pilihan tidak valid")
		InputData()
	}
}

func InputTim(a *TabTim) {
	var b TabPemain
	fmt.Println("=== INPUT DATA TIM ===")
	for i := 0; i < 6; i++ {
		fmt.Printf("\nData Tim ke-%d:\n", i+1)
		fmt.Println("Masukkan nama tim: ")
		fmt.Scan(&a[i].Nama)
		InputPemain(&b)
	}

	MenuAdmin()

}

func InputPemain(a *TabPemain) {
	fmt.Println("Masukkan nama pemain dan posisi: ")
	for i := 0; i < 7; i++ {
		fmt.Scanln(&a[i].Nama, &a[i].Posisi)
	}
}

func InputPertandingan(a *TabPertandingan) {
	var d, e TabTim
	fmt.Println("=== INPUT JADWAL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("\nData Pertandingan ke-%d:\n", i+1)
		fmt.Print("Masukkan nama tim 1: ")
		fmt.Scan(&a[i].Tim1)
		fmt.Print("Masukkan nama tim 2: ")
		fmt.Scan(&a[i].Tim2)
		b := Validtim(d, &a[i].Tim1)
		c := Validtim(e, &a[i].Tim2)
		if !b || !c {
			fmt.Println("Tim tidak valid, silakan masukkan nama tim yang benar.")
			i--
		}
		fmt.Print("Masukkan hari pertandingan: ")
		fmt.Scan(&a[i].Hari)
	}
	MenuAdmin()
}

func Validtim(a TabTim, nama *string) bool {
	for i := 0; i < 6; i++ {
		if a[i].Nama == *nama {
			return true
		}
	}
	return false
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
	*n = *n - 1
	fmt.Println("Data berhasil dihapus.")
	MenuAdmin()
}

func InputHasil(a TabPertandingan) {
	var b, c string
	fmt.Println("=== INPUT HASIL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Masukkan hasil pertandingan %d:\n", i+1)
		fmt.Printf("Tim 1 (%s): ", a[i].Tim1)
		fmt.Scan(&b)
		fmt.Printf("Tim 2 (%s): ", a[i].Tim2)
		fmt.Scan(&c)
		a[i].HasilTim1 = b
		a[i].HasilTim2 = c
	}
	MenuAdmin()
}

func tampilData() {
	var pilihan int
	fmt.Println("---MENU INPUT---")
	fmt.Println("1. Tampil Data Tim")
	fmt.Println("2. Tampil jadwal")
	fmt.Println("3. Tampil skor pertandingan")
	fmt.Println("4. Kembali ke Menu Admin")
	fmt.Scan(&pilihan)

	var tim TabTim
	var pemain [7]TabPemain
	var pertandingan TabPertandingan

	switch pilihan {
	case 1:
		TampilTim(tim, pemain)
	case 2:
		TampilPertandingan(pertandingan)
	case 3:
		TampilHasil(pertandingan)
	case 4:
		MenuAdmin()
	default:
		fmt.Println("Pilihan tidak valid")
		InputData()
	}
}

func TampilTim(a TabTim, b [7]TabPemain) {
	fmt.Println("=== DATA TIM ===")
	for i := 0; i < 6; i++ {
		fmt.Printf("%-3d | %-15s | %-5d | %-5d | %-5d | %-5.2f\n", i+1, a[i].Nama, a[i].JumlahPertandingan, a[i].JumlahMenang, a[i].JumlahKalah, a[i].RasioMenang)
		fmt.Println("   Daftar Pemain:")
		for j := 0; j < 7; j++ {
			fmt.Printf("      - %s (%s)\n", b[i][j].Nama, b[i][j].Posisi)
		}
	}
	tampilData()
}

func TampilPertandingan(a TabPertandingan) {
	fmt.Println("=== JADWAL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s vs %s, Hari: %s\n", i+1, a[i].Tim1, a[i].Tim2, a[i].Hari)
	}
	tampilData()
}

func TampilHasil(a TabPertandingan) {
	fmt.Println("=== HASIL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s (%s) vs %s (%s)\n", i+1, a[i].Tim1, a[i].HasilTim1, a[i].Tim2, a[i].HasilTim2)
	}
	tampilData()
}

func Peringkat(a *TabTim) {
	fmt.Println("=== PERINGKAT TIM ===")
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 7; j++ {
			if a[i].RasioMenang < a[j].RasioMenang {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for i := 0; i < 7; i++ {
		fmt.Printf("%d. %s - Rasio Menang: %.2f\n", i+1, a[i].Nama, a[i].RasioMenang)
	}
	tampilData()
}

func MenuPengguna() {
	var pilihan int
	var pemain [7]TabPemain
	var tim TabTim

	fmt.Println("---MENU PENGGUNA---")
	fmt.Println("1. Tampilkan Data Tim")
	fmt.Println("2. Tampilkan Jadwal Pertandingan")
	fmt.Println("3. Tampilkan Hasil Pertandingan")
	fmt.Println("4. Tampilkan Peringkat Tim")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tampilTim2(tim, pemain)
	case 2:
		TampilPertandingan2(TabPertandingan{})
	case 3:
		TampilHasil2(TabPertandingan{})
	case 4:
		Peringkat(&tim)
	case 5:
		Menu()
	default:
		fmt.Println("Pilihan tidak valid")
		MenuPengguna()
	}
}

func tampilTim2(a TabTim, b [7]TabPemain) {
	fmt.Println("=== DATA TIM ===")
	for i := 0; i < 6; i++ {
		fmt.Printf("%-3d | %-15s | %-5d | %-5d | %-5d | %-5.2f\n", i+1, a[i].Nama, a[i].JumlahPertandingan, a[i].JumlahMenang, a[i].JumlahKalah, a[i].RasioMenang)
		fmt.Println("   Daftar Pemain:")
		for j := 0; j < 7; j++ {
			fmt.Printf("      - %s (%s)\n", b[i][j].Nama, b[i][j].Posisi)
		}
	}
	MenuPengguna()
}

func TampilPertandingan2(a TabPertandingan) {
	fmt.Println("=== JADWAL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s vs %s, Hari: %s\n", i+1, a[i].Tim1, a[i].Tim2, a[i].Hari)
	}
	MenuPengguna()
}

func TampilHasil2(a TabPertandingan) {
	fmt.Println("=== HASIL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s (%s) vs %s (%s)\n", i+1, a[i].Tim1, a[i].HasilTim1, a[i].Tim2, a[i].HasilTim2)
	}
	MenuPengguna()
}

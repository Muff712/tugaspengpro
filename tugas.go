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

type TabArrayPemain [6]TabPemain

type Pertandingan struct {
	Tim1, Tim2 string
	Hari       string
	HasilTim1  string
	HasilTim2  string
}

type TabPertandingan [9]Pertandingan

var tim TabTim
var pemain TabArrayPemain
var pertandingan TabPertandingan

func main() {
	Menu()
}

func Menu() {
	var pilihan string
	fmt.Println("Pilih pengguna:")
	fmt.Println(" Admin")
	fmt.Println(" User")
	fmt.Println(" Keluar")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case "admin", "Admin":
		MenuAdmin()
	case "user", "User":
		MenuPengguna()
	case "keluar", "Keluar":
		fmt.Println("Anda keluar")
		return
	default:
		fmt.Println("Pilihan tidak valid")
		Menu()
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
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		InputData()
	case 2:
		HapusData()
	case 3:
		tampilData()
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
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		InputTim()
	case 2:
		InputPertandingan()
	case 3:
		InputHasil()
	case 4:
		MenuAdmin()
	default:
		fmt.Println("Pilihan tidak valid")
		InputData()
	}
}

func InputTim() {
	fmt.Println("=== INPUT DATA TIM ===")
	for i := 0; i < 6; i++ {
		fmt.Printf("\nData Tim ke-%d:\n", i+1)
		fmt.Print("Masukkan nama tim: ")
		fmt.Scan(&tim[i].Nama)
		InputPemain(&pemain[i])
	}
	MenuAdmin()
}

func InputPemain(p *TabPemain) {
	fmt.Println("Masukkan nama pemain dan posisi:")
	for i := 0; i < 7; i++ {
		fmt.Printf("Pemain %d - Nama dan Posisi: ", i+1)
		fmt.Scan(&p[i].Nama, &p[i].Posisi)
	}
}

func InputPertandingan() {
	fmt.Println("=== INPUT JADWAL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("\nData Pertandingan ke-%d:\n", i+1)
		fmt.Print("Masukkan nama tim 1: ")
		fmt.Scan(&pertandingan[i].Tim1)
		fmt.Print("Masukkan nama tim 2: ")
		fmt.Scan(&pertandingan[i].Tim2)
		if !Validtim(pertandingan[i].Tim1) || !Validtim(pertandingan[i].Tim2) {
			fmt.Println("Tim tidak valid, silakan masukkan nama tim yang benar.")
			i--
			continue
		}
		fmt.Print("Masukkan hari pertandingan: ")
		fmt.Scan(&pertandingan[i].Hari)
	}
	MenuAdmin()
}

func Validtim(nama string) bool {
	for i := 0; i < 6; i++ {
		if tim[i].Nama == nama {
			return true
		}
	}
	return false
}

func CariIdxTim(nama string) int {
	for i := 0; i < 6; i++ {
		if tim[i].Nama == nama {
			return i
		}
	}
	return -1
}

func HapusData() {
	var nama string
	fmt.Println("=== HAPUS TIM ===")
	fmt.Print("Masukkan nama tim yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx := CariIdxTim(nama)
	if idx == -1 {
		fmt.Println("Tim tidak ditemukan.")
		return
	}
	for i := idx; i < 5; i++ {
		tim[i] = tim[i+1]
		pemain[i] = pemain[i+1]
	}
	tim[5] = Tim{}
	pemain[5] = TabPemain{}
	fmt.Println("Data berhasil dihapus.")
	MenuAdmin()
}

func InputHasil() {
	fmt.Println("=== INPUT HASIL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d - %s vs %s:\n", i+1, pertandingan[i].Tim1, pertandingan[i].Tim2)
		fmt.Printf("Skor %s: ", pertandingan[i].Tim1)
		fmt.Scan(&pertandingan[i].HasilTim1)
		fmt.Printf("Skor %s: ", pertandingan[i].Tim2)
		fmt.Scan(&pertandingan[i].HasilTim2)
	}
	MenuAdmin()
}

func tampilData() {
	var pilihan int
	fmt.Println("---MENU TAMPIL---")
	fmt.Println("1. Tampil Data Tim")
	fmt.Println("2. Tampil Jadwal")
	fmt.Println("3. Tampil Skor Pertandingan")
	fmt.Println("4. Kembali ke Menu Admin")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		TampilTim()
	case 2:
		TampilPertandingan()
	case 3:
		TampilHasil()
	case 4:
		MenuAdmin()
	default:
		fmt.Println("Pilihan tidak valid")
		tampilData()
	}
}

func TampilTim() {
	fmt.Println("=== DATA TIM ===")
	for i := 0; i < 6; i++ {
		if tim[i].Nama != "" {
			fmt.Printf("%d. %s\n", i+1, tim[i].Nama)
			for j := 0; j < 7; j++ {
				if pemain[i][j].Nama != "" {
					fmt.Printf("    - %s (%s)\n", pemain[i][j].Nama, pemain[i][j].Posisi)
				}
			}
		}
	}
	tampilData()
}

func TampilPertandingan() {
	fmt.Println("=== JADWAL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s vs %s, Hari: %s\n", i+1, pertandingan[i].Tim1, pertandingan[i].Tim2, pertandingan[i].Hari)
	}
	tampilData()
}

func TampilHasil() {
	fmt.Println("=== HASIL PERTANDINGAN ===")
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d: %s (%s) vs %s (%s)\n", i+1, pertandingan[i].Tim1, pertandingan[i].HasilTim1, pertandingan[i].Tim2, pertandingan[i].HasilTim2)
	}
	tampilData()
}

func MenuPengguna() {
	var pilihan int
	fmt.Println("---MENU PENGGUNA---")
	fmt.Println("1. Tampilkan Data Tim")
	fmt.Println("2. Tampilkan Jadwal Pertandingan")
	fmt.Println("3. Tampilkan Hasil Pertandingan")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tampilTim2(tim, pemain)
	case 2:
		TampilPertandingan()
	case 3:
		TampilHasil()
	case 4:
		Menu()
	default:
		fmt.Println("Pilihan tidak valid")
		MenuPengguna()
	}
}

func PeringkatTim(a TabTim) {

}

func tampilTim2(a TabTim, b TabArrayPemain) {
	fmt.Println("=== DATA TIM ===")
	dataDitemukan := false

	for i := 0; i < 6; i++ {
		if a[i].Nama != "" {
			dataDitemukan = true
			fmt.Printf("%-3d | %-15s | %-5d | %-5d | %-5d | %-5.2f\n",
				i+1, a[i].Nama, a[i].JumlahPertandingan, a[i].JumlahMenang, a[i].JumlahKalah, a[i].RasioMenang)
			fmt.Println("  Daftar Pemain:")
			for j := 0; j < 7; j++ {
				if b[i][j].Nama != "" {
					fmt.Printf("      - %s (%s)\n", b[i][j].Nama, b[i][j].Posisi)
				}
			}
		}
	}

	if !dataDitemukan {
		fmt.Println("Belum ada data tim yang tersedia.")
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

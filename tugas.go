package main

import "fmt"


type Tim struct {
	Nama                                                                                string
	JumlahPertandingan,  JumlahMenang,  JumlahKalah int
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
	HasilTim1   string
	HasilTim2   string
}

type TabPertandingan [9 - 1]Pertandingan



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
	}
	fmt.Println("Pilihan tidak valid")
	InputData()
}

func InputTim(a *TabTim) {
	var b TabPemain
	fmt.Println("=== INPUT DATA TIM ===")
	for i := 0; i < 7; i++ {
		fmt.Printf("\nData Tim ke-%d:\n", i+1)
		fmt.Println("Masukkan nama tim: ")
		fmt.Scan(&a[i].Nama)
		InputPemain(&b)
		}
	}

func InputPemain(a *TabPemain) {
	for i := 0;i < 9; i++ {
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

    n := 8
    i := 0
    for i < n {
        x := i + 1
        for x < n {
            samaTim := (a[i].Tim1 == a[x].Tim1 && a[i].Tim2 == a[x].Tim2) ||
                       (a[i].Tim1 == a[x].Tim2 && a[i].Tim2 == a[x].Tim1)
            if samaTim && a[i].Hari == a[x].Hari {
                for y := x; y < n-1; y++ {
                    a[y] = a[y+1]
                }
                n--
            } else {
                x++
            }
        }
        i++
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
    fmt.Println("Data berhasil dihapus.")
}

func InputHasil(a TabPertandingan){
	var b ,c string
	var d TabTim
	fmt.Println("=== INPUT HASIL PERTANDINGAN ===")
	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan hasil pertandingan %d:\n", i+1)
		fmt.Printf("Tim 1 (%s): ", a[i].Tim1)
		fmt.Scan(&b)
		fmt.Printf("Tim 2 (%s): ", a[i].Tim2)
		fmt.Scan(&c)
		a[i].HasilTim1 = b
		a[i].HasilTim2 = c
	}
	for i := 0; i < 8; i++ {
		if a[i].Tim1 == d[i].Nama {
			d[i].JumlahPertandingan++
			if b == "Menang" || b == "menang" {
				d[i].JumlahMenang++
			} else if b == "Kalah" || b == "kalah" {
				d[i].JumlahKalah++
			}
			d[i].RasioMenang = float32(d[i].JumlahMenang) / float32(d[i].JumlahPertandingan)
		}
	}
	for i := 0; i < 8; i++ {
		if a[i].Tim2 == d[i].Nama {
			d[i].JumlahPertandingan++
			if c == "Menang" || c == "menang" {
				d[i].JumlahMenang++
			} else if c == "Kalah" || c == "kalah" {
				d[i].JumlahKalah++
			}
			d[i].RasioMenang = float32(d[i].JumlahMenang) / float32(d[i].JumlahPertandingan)
		}
	}

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
	var pemain [10]TabPemain
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
	}
	fmt.Println("Pilihan tidak valid")
	InputData()
}

func TampilTim(a TabTim, b [10]TabPemain) {
    fmt.Println("=== DATA TIM ===")
    for i := 0; i < 9; i++ {
		fmt.Printf("%-3d | %-15s | %-5d | %-5d | %-5d | %-5.2f | %-8.2f\n", i+1, a[i].Nama, a[i].JumlahPertandingan, a[i].JumlahMenang, a[i].JumlahKalah, a[i].RasioMenang, a[i].RasioMenang)
        fmt.Println("   Daftar Pemain:")
        for j := 0; j < 9; j++ {
            fmt.Printf("      - %s (%s)\n", b[i][j].Nama, b[i][j].Posisi)
        }
    }
}

func TampilPertandingan(a TabPertandingan) {
	fmt.Println("=== JADWAL PERTANDINGAN ===")
	for i := 0; i < 8; i++ {
		fmt.Printf("Pertandingan %d: %s vs %s, Hari: %s\n", i+1, a[i].Tim1, a[i].Tim2, a[i].Hari)
	}
}

func TampilHasil(a TabPertandingan) {
	fmt.Println("=== HASIL PERTANDINGAN ===")
	for i := 0; i < 8; i++ {
		fmt.Printf("Pertandingan %d: %s (%s) vs %s (%s)\n", i+1, a[i].Tim1, a[i].HasilTim1, a[i].Tim2, a[i].HasilTim2)
	}
}

func Peringkat(a *TabTim) {
	fmt.Println("=== PERINGKAT TIM ===")
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 7; j++ {
			if a[i].RasioMenang < a[j].RasioMenang {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for i := 0; i < 7; i++ {
		fmt.Printf("%d. %s - Rasio Menang: %.2f\n", i+1, a[i].Nama, a[i].RasioMenang)
	}
}

func MenuPengguna() {
	var pilihan int
	fmt.Println("---MENU PENGGUNA---")
	fmt.Println("1. Tampilkan Data Tim")
	fmt.Println("2. Tampilkan Jadwal Pertandingan")
	fmt.Println("3. Tampilkan Hasil Pertandingan")
	fmt.Println("4. Tampilkan Peringkat Tim")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Scan(&pilihan)

	var pertandingan TabPertandingan
	var tim TabTim
	var pemain [10]TabPemain
	switch pilihan {
	case 1:
		TampilTim(tim, pemain)
	case 2:
		TampilPertandingan(pertandingan)
	case 3:
		TampilHasil(pertandingan)
	case 4:
		Peringkat(&tim)
	case 5:
		Menu()
	default:
		fmt.Println("Pilihan tidak valid")
		MenuPengguna()
	}
}









package main

import "fmt"

const (
	MaksimalTim          = 100
	MaksimalPemain       = 100
	MaksimalPertandingan = 100
)

type Tim struct {
	Nama                                                                                string
	JumlahPertandingan, Peringkat, JumlahMenang, JumlahSeri, JumlahKalah, JumlahAnggota int
	RasioMenang                                                                         float32
}

type TabTim [MaksimalTim - 1]Tim


type Pemain struct {
	Nama               string
	Posisi             string
}

type TabPemain [MaksimalPemain - 1]Pemain


type Pertandingan struct {
	Tim1, Tim2  string
	Hari        string
	HasilTim1   int
	HasilTim2   int
}

type TabPertandingan [MaksimalPertandingan - 1]Pertandingan


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
	fmt.Println("2. Input Data Pemain")
	fmt.Println("3. Input Data Pertandingan")
	fmt.Println("4. Kembali ke Menu Admin")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		InputTim(daftarTim)
	case 2:
		InputPemain(daftarPemain, daftarTim)
	case 3:
		InputPertandingan(daftarPertandingan, daftarTim)
	case 4:
		MenuAdmin()
	}
	fmt.Println("Pilihan tidak valid")
	InputData()
}

func InputTim(a *TabTim, n int) {
	fmt.Print("Masukkan jumlah tim: ")
	fmt.Scan(&n)

	if n <= 0 || n > MaksimalTim {
		fmt.Printf("Jumlah tim harus antara 1 dan %d\n", MaksimalTim)
		return
	}

	for i := 0; i < n; i++ {
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

func InputPertandingan(daftarPertandingan *DaftarPertandingan, daftarTim *DaftarTim) {
	var jumlah int
	fmt.Print("Masukkan jumlah pertandingan: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 || jumlah > MaksimalPertandingan {
		fmt.Printf("Jumlah pertandingan harus antara 1 dan %d\n", MaksimalPertandingan)
		return
	}

	i := 0
	for i < jumlah {
		fmt.Printf("\nData Pertandingan ke-%d:\n", i+1)
		var pertandingan Pertandingan
		var inputValid bool
		inputValid = true

		fmt.Print("Masukkan nama tim 1: ")
		fmt.Scan(&pertandingan.Tim1)
		fmt.Print("Masukkan nama tim 2: ")
		fmt.Scan(&pertandingan.Tim2)

		if pertandingan.Tim1 == pertandingan.Tim2 {
			fmt.Println("Tim tidak boleh sama")
			inputValid = false
		}

		if inputValid {
			fmt.Print("Masukkan hari (senin-minggu): ")
			fmt.Scan(&pertandingan.Hari)
			fmt.Print("Masukkan jam (0-23): ")
			fmt.Scan(&pertandingan.Jam)

			if pertandingan.Jam < 0 || pertandingan.Jam > 23 {
				fmt.Println("Jam tidak valid")
				inputValid = false
			}

			if inputValid {
				daftarPertandingan.Pertandingan[i] = pertandingan
				daftarPertandingan.Jumlah++
				i++
			}
		}
	}
}

func TampilkanData(daftarTim *TabTim, daftarPemain *TabPemain, daftarPertandingan *TabPertandingan) {
	fmt.Println("\n=== DATA TIM ===")
	for i := 0; i < daftarTim.Jumlah; i++ {
		fmt.Printf("\nTim: %s\n", daftarTim.Tim[i].Nama)
		fmt.Printf("Jumlah Pertandingan: %d\n", daftarTim.Tim[i].JumlahPertandingan)
		fmt.Printf("Statistik: Menang:%d Seri:%d Kalah:%d\n",
			daftarTim.Tim[i].JumlahMenang,
			daftarTim.Tim[i].JumlahSeri,
			daftarTim.Tim[i].JumlahKalah)
		fmt.Printf("Rasio Kemenangan: %.2f%%\n", daftarTim.Tim[i].RasioMenang)
	}

	fmt.Println("\n=== DATA PEMAIN ===")
	for i := 0; i < daftarPemain.Jumlah; i++ {
		fmt.Printf("\nNama: %s\n", daftarPemain.Pemain[i].Nama)
		fmt.Printf("Tim: %s\n", daftarPemain.Pemain[i].NamaTim)
		fmt.Printf("Posisi: %s\n", daftarPemain.Pemain[i].Posisi)
	}

	fmt.Println("\n=== JADWAL PERTANDINGAN ===")
	for i := 0; i < daftarPertandingan.Jumlah; i++ {
		fmt.Printf("\n%s vs %s\n",
			daftarPertandingan.Pertandingan[i].Tim1,
			daftarPertandingan.Pertandingan[i].Tim2)
		fmt.Printf("Jadwal: %s jam %d:00\n",
			daftarPertandingan.Pertandingan[i].Hari,
			daftarPertandingan.Pertandingan[i].Jam)
	}
}

func HapusData(daftarTim *DaftarTim, daftarPemain *DaftarPemain, daftarPertandingan *DaftarPertandingan) {
	var pilihan int
	fmt.Println("1. Hapus Tim")
	fmt.Println("2. Hapus Pemain")
	fmt.Println("3. Hapus Pertandingan")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Masukkan nama tim yang akan dihapus: ")
		fmt.Scan(&nama)

		for i := 0; i < daftarTim.Jumlah; i++ {
			if daftarTim.Tim[i].Nama == nama {
				for j := i; j < daftarTim.Jumlah-1; j++ {
					daftarTim.Tim[j] = daftarTim.Tim[j+1]
				}
				daftarTim.Jumlah--
				fmt.Println("Tim berhasil dihapus")
				return
			}
		}
		fmt.Println("Tim tidak ditemukan")

	case 2:
		var nama string
		fmt.Print("Masukkan nama pemain yang akan dihapus: ")
		fmt.Scan(&nama)

		for i := 0; i < daftarPemain.Jumlah; i++ {
			if daftarPemain.Pemain[i].Nama == nama {
				for j := i; j < daftarPemain.Jumlah-1; j++ {
					daftarPemain.Pemain[j] = daftarPemain.Pemain[j+1]
				}
				daftarPemain.Jumlah--
				fmt.Println("Pemain berhasil dihapus")
				return
			}
		}
		fmt.Println("Pemain tidak ditemukan")

	case 3:
		var tim1, tim2 string
		fmt.Print("Masukkan nama tim 1: ")
		fmt.Scan(&tim1)
		fmt.Print("Masukkan nama tim 2: ")
		fmt.Scan(&tim2)

		for i := 0; i < daftarPertandingan.Jumlah; i++ {
			if (daftarPertandingan.Pertandingan[i].Tim1 == tim1 &&
				daftarPertandingan.Pertandingan[i].Tim2 == tim2) {
				for j := i; j < daftarPertandingan.Jumlah-1; j++ {
					daftarPertandingan.Pertandingan[j] = daftarPertandingan.Pertandingan[j+1]
				}
				daftarPertandingan.Jumlah--
				fmt.Println("Pertandingan berhasil dihapus")
				return
			}
		}
		fmt.Println("Pertandingan tidak ditemukan")
	}
}

func MenuPengguna() {
	// Implementasi menu untuk pengguna biasa
}

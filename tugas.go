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

type DaftarTim struct {
	Tim []Tim
}

type Pemain struct {
	Nama               string
	NamaTim            string
	Posisi             string
}

type DaftarPemain struct {
	Pemain []Pemain
	Jumlah int
}

type Pertandingan struct {
	Tim1, Tim2  string
	Hari        string
	Jam         int
	HasilTim1   int
	HasilTim2   int
}

type DaftarPertandingan struct {
	Pertandingan []Pertandingan
	Jumlah       int
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
	fmt.Println("4. Keluar")
	fmt.Println("5. Kembali ke Menu Admin")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		daftarTim := &DaftarTim{Tim: make([]Tim, MaksimalTim)}
		InputTim(daftarTim)
	case 2:
		daftarPemain := &DaftarPemain{Pemain: make([]Pemain, MaksimalPemain)}
		InputPemain(daftarPemain, daftarTim)
	case 3:
		daftarPertandingan := &DaftarPertandingan{Pertandingan: make([]Pertandingan, MaksimalPertandingan)}
		InputPertandingan(daftarPertandingan, daftarTim)
	case 4:
		return
	case 5:
		MenuAdmin()
	}
	fmt.Println("Pilihan tidak valid")
	InputData()
}

func InputTim(daftarTim *DaftarTim) {
	var jumlah int
	fmt.Print("Masukkan jumlah tim: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 || jumlah > MaksimalTim {
		fmt.Printf("Jumlah tim harus antara 1 dan %d\n", MaksimalTim)
		return
	}

	i := 0
	for i < jumlah {
		fmt.Printf("\nData Tim ke-%d:\n", i+1)
		var tim Tim
		var inputValid bool
		inputValid = true

		fmt.Print("Masukkan nama tim: ")
		fmt.Scan(&tim.Nama)

		// Cek duplikasi nama tim
		for j := 0; j < i; j++ {
			if daftarTim.Tim[j].Nama == tim.Nama {
				fmt.Printf("Tim dengan nama %s sudah ada\n", tim.Nama)
				inputValid = false
			}
		}

		if inputValid {
			var adaRiwayat string
			fmt.Print("Ada riwayat pertandingan? (ya/tidak): ")
			fmt.Scan(&adaRiwayat)

			if adaRiwayat == "ya" || adaRiwayat == "Ya" {
				fmt.Print("Masukkan jumlah pertandingan: ")
				fmt.Scan(&tim.JumlahPertandingan)

				if tim.JumlahPertandingan > 0 {
					fmt.Print("Masukkan jumlah menang: ")
					fmt.Scan(&tim.JumlahMenang)
					fmt.Print("Masukkan jumlah seri: ")
					fmt.Scan(&tim.JumlahSeri)
					fmt.Print("Masukkan jumlah kalah: ")
					fmt.Scan(&tim.JumlahKalah)

					totalPertandingan := tim.JumlahMenang + tim.JumlahSeri + tim.JumlahKalah
					if totalPertandingan == tim.JumlahPertandingan {
						tim.RasioMenang = float32(tim.JumlahMenang) / float32(tim.JumlahPertandingan) * 100
						daftarTim.Tim[i] = tim
						i++
					} else {
						fmt.Println("Total pertandingan tidak sesuai dengan jumlah menang, seri, dan kalah")
						inputValid = false
					}
				} else {
					fmt.Println("Jumlah pertandingan harus lebih dari 0")
					inputValid = false
				}
			} else {
				daftarTim.Tim[i] = tim
				i++
			}
		}
	}
}

func InputPemain(daftarPemain *DaftarPemain, daftarTim *DaftarTim) {
	var jumlah int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 || jumlah > MaksimalPemain {
		fmt.Printf("Jumlah pemain harus antara 1 dan %d\n", MaksimalPemain)
		return
	}

	i := 0
	for i < jumlah {
		fmt.Printf("\nData Pemain ke-%d:\n", i+1)
		var pemain Pemain
		var inputValid bool
		inputValid = true

		fmt.Print("Masukkan nama pemain: ")
		fmt.Scan(&pemain.Nama)
		fmt.Print("Masukkan nama tim: ")
		fmt.Scan(&pemain.NamaTim)

		timDitemukan := false
		for j := 0; j < daftarTim.Jumlah; j++ {
			if daftarTim.Tim[j].Nama == pemain.NamaTim {
				timDitemukan = true
			}
		}

		if !timDitemukan {
			fmt.Printf("Tim %s tidak ditemukan\n", pemain.NamaTim)
			inputValid = false
		}

		if inputValid {
			fmt.Print("Masukkan posisi pemain: ")
			fmt.Scan(&pemain.Posisi)

			pemain.JumlahPertandingan = 0
			pemain.JumlahMenang = 0
			pemain.JumlahSeri = 0
			pemain.JumlahKalah = 0

			daftarPemain.Pemain[i] = pemain
			daftarPemain.Jumlah++
			i++
		}
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

func TampilkanData(daftarTim *DaftarTim, daftarPemain *DaftarPemain, daftarPertandingan *DaftarPertandingan) {
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

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

type Team struct {
	Name                string
	MatchesPlayed       int
	Ranking             int
	Wins, Draws, Losses int
	MemberCount         int
	TotalTeams          int
	MatchRate           float32
}

type Player struct {
	Name, TeamName      string
	MatchesPlayed       int
	Ranking             int
	Wins, Draws, Losses int
	Position            string
	MemberCount         int
}

type Match struct {
	Team1, Team2 string
	Day          string
	Hour         int
}

type TeamArray [MaxTeams]Team

type PlayerArray [MaxPlayers]Player

type MatchArray [MaxMatches]Match

var (
	totalTeams, totalPlayers, totalMatches int

	teams   TeamArray
	players PlayerArray
	matches MatchArray
)

func main() {
	for {
		if err := mainMenu(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func mainMenu() error {
	fmt.Println("=== PILIH PENGGUNA ===")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Println("3. Keluar")

	var choice int
	fmt.Print("Pilihan: ")
	if _, err := fmt.Scan(&choice); err != nil {
		return fmt.Errorf("input tidak valid: %v", err)
	}

	switch choice {
	case 1:
		return adminMenu()
	case 2:
		return userMenu()
	case 3:
		fmt.Println("Anda keluar")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
}

func adminMenu() error {
	fmt.Println("\n=== MENU ADMIN ===")
	fmt.Println("1. Input Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Tampilkan Data")
	fmt.Println("4. Keluar")

	var choice int
	fmt.Print("Pilihan: ")
	if _, err := fmt.Scan(&choice); err != nil {
		return fmt.Errorf("input tidak valid: %v", err)
	}

	switch choice {
	case 1:
		return inputDataMenu()
	case 2:
		return deleteDataMenu()
	case 3:
		return userMenu()
	case 4:
		fmt.Println("Kembali ke menu utama")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
}

func userMenu() error {
	fmt.Println("\n=== MENU USER ===")
	fmt.Println("1. Tampilkan Data Tim")
	fmt.Println("2. Tampilkan Data Pemain")
	fmt.Println("3. Tampilkan Data Pertandingan")
	fmt.Println("4. Tampilkan Tim dengan Winrate Tertinggi")
	fmt.Println("5. Tampilkan Tim dengan Winrate Terendah")
	fmt.Println("6. Keluar")

	var choice int
	fmt.Print("Pilihan: ")
	if _, err := fmt.Scan(&choice); err != nil {
		return fmt.Errorf("input tidak valid: %v", err)
	}

	switch choice {
	case 1:
		tampilTim()
	case 2:
		tampilPemain()
	case 3:
		tampilPertandingan()
	case 4:
		tampilWinrateTertinggi()
	case 5:
		tampilWinrateTerendah()
	case 6:
		fmt.Println("Kembali ke menu utama")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
	return nil
}

func inputDataMenu() error {
	fmt.Println("\n=== MENU INPUT DATA ===")
	fmt.Println("1. Input Data Tim")
	fmt.Println("2. Input Data Pemain")
	fmt.Println("3. Input Data Pertandingan")
	fmt.Println("4. Kembali")

	var choice int
	fmt.Print("Pilihan: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		inputTeam()
	case 2:
		inputPlayer()
	case 3:
		inputMatch()
	case 4:
		fmt.Println("Kembali ke menu admin")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
	return nil
}

func deleteDataMenu() error {
	fmt.Println("\n=== MENU HAPUS DATA ===")
	fmt.Println("1. Hapus Data Tim")
	fmt.Println("2. Hapus Data Pemain")
	fmt.Println("3. Kembali")

	var choice int
	fmt.Print("Pilihan: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		deleteTeam()
	case 2:
		deletePlayer()
	case 3:
		fmt.Println("Kembali ke menu admin")
		return nil
	default:
		return fmt.Errorf("pilihan tidak valid")
	}
	return nil
}

func inputTeam() {
	if totalTeams >= MaxTeams {
		fmt.Println("Kapasitas tim penuh")
		return
	}
	fmt.Print("Nama Tim: ")
	fmt.Scan(&teams[totalTeams].Name)
	totalTeams++
	fmt.Println("Data tim berhasil ditambahkan.")
}

func inputPlayer() {
	if totalPlayers >= MaxPlayers {
		fmt.Println("Kapasitas pemain penuh")
		return
	}
	fmt.Print("Nama Pemain: ")
	fmt.Scan(&players[totalPlayers].Name)
	fmt.Print("Nama Tim: ")
	fmt.Scan(&players[totalPlayers].TeamName)
	totalPlayers++
	fmt.Println("Data pemain berhasil ditambahkan.")
}

func inputMatch() {
	if totalMatches >= MaxMatches {
		fmt.Println("Kapasitas pertandingan penuh")
		return
	}
	var team1, team2, day, result string
	var hour int

	fmt.Print("Tim 1: ")
	fmt.Scan(&team1)
	fmt.Print("Tim 2: ")
	fmt.Scan(&team2)
	fmt.Print("Hari: ")
	fmt.Scan(&day)
	fmt.Print("Jam (0-23): ")
	fmt.Scan(&hour)
	fmt.Print("Hasil (1=Tim1 menang, 2=Tim2 menang, D=Draw): ")
	fmt.Scan(&result)

	matches[totalMatches] = Match{Team1: team1, Team2: team2, Day: day, Hour: hour}
	totalMatches++

	updateStats(team1, team2, result)

	fmt.Println("Data pertandingan dan statistik berhasil ditambahkan.")
}

func updateStats(team1, team2, result string) {
	idx1 := findTeamIndex(team1)
	idx2 := findTeamIndex(team2)
	if idx1 == -1 || idx2 == -1 {
		fmt.Println("Nama tim tidak ditemukan.")
		return
	}

	teams[idx1].MatchesPlayed++
	teams[idx2].MatchesPlayed++

	switch strings.ToUpper(result) {
	case "1":
		teams[idx1].Wins++
		teams[idx2].Losses++
	case "2":
		teams[idx2].Wins++
		teams[idx1].Losses++
	case "D":
		teams[idx1].Draws++
		teams[idx2].Draws++
	default:
		fmt.Println("Hasil tidak valid")
		return
	}

	teams[idx1].MatchRate = float32(teams[idx1].Wins) / float32(teams[idx1].MatchesPlayed) * 100
	teams[idx2].MatchRate = float32(teams[idx2].Wins) / float32(teams[idx2].MatchesPlayed) * 100
}

func findTeamIndex(name string) int {
	for i := 0; i < totalTeams; i++ {
		if strings.EqualFold(teams[i].Name, name) {
			return i
		}
	}
	return -1
}

func tampilTim() {
	fmt.Println("\n=== DATA TIM ===")
	for i := 0; i < totalTeams; i++ {
		t := teams[i]
		fmt.Printf("%d. %s | Main: %d | W: %d | D: %d | L: %d | Winrate: %.2f%%\n",
			i+1, t.Name, t.MatchesPlayed, t.Wins, t.Draws, t.Losses, t.MatchRate)
	}
}

func tampilPemain() {
	fmt.Println("\n=== DATA PEMAIN ===")
	for i := 0; i < totalPlayers; i++ {
		fmt.Printf("%d. %s - %s\n", i+1, players[i].Name, players[i].TeamName)
	}
}

func tampilPertandingan() {
	fmt.Println("\n=== DATA PERTANDINGAN ===")
	for i := 0; i < totalMatches; i++ {
		fmt.Printf("%d. %s vs %s - %s %d:00\n", i+1, matches[i].Team1, matches[i].Team2, matches[i].Day, matches[i].Hour)
	}
}

func tampilWinrateTertinggi() {
	if totalTeams == 0 {
		fmt.Println("Belum ada data tim.")
		return
	}
	maxIdx := 0
	for i := 1; i < totalTeams; i++ {
		if teams[i].MatchRate > teams[maxIdx].MatchRate {
			maxIdx = i
		}
	}
	t := teams[maxIdx]
	fmt.Printf("Tim dengan winrate tertinggi: %s (%.2f%%)\n", t.Name, t.MatchRate)
}

func tampilWinrateTerendah() {
	if totalTeams == 0 {
		fmt.Println("Belum ada data tim.")
		return
	}
	minIdx := 0
	for i := 1; i < totalTeams; i++ {
		if teams[i].MatchRate < teams[minIdx].MatchRate {
			minIdx = i
		}
	}
	t := teams[minIdx]
	fmt.Printf("Tim dengan winrate terendah: %s (%.2f%%)\n", t.Name, t.MatchRate)
}

func deleteTeam() {
	var name string
	fmt.Print("Masukkan nama tim yang akan dihapus: ")
	fmt.Scan(&name)
	for i := 0; i < totalTeams; i++ {
		if strings.EqualFold(teams[i].Name, name) {
			teams[i] = teams[totalTeams-1]
			totalTeams--
			fmt.Println("Tim berhasil dihapus.")
			return
		}
	}
	fmt.Println("Tim tidak ditemukan.")
}

func deletePlayer() {
	var name string
	fmt.Print("Masukkan nama pemain yang akan dihapus: ")
	fmt.Scan(&name)
	for i := 0; i < totalPlayers; i++ {
		if strings.EqualFold(players[i].Name, name) {
			players[i] = players[totalPlayers-1]
			totalPlayers--
			fmt.Println("Pemain berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pemain tidak ditemukan.")
}

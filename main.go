package main

import (
	"fmt"
	"os"
	"strconv"
)

// Define the Friend struct to represent the attributes of a friend
type Friend struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Define a slice of Friend to store friend data
var friends = []Friend{
	{1, "Deni Romadhon", "Jakarta", "Karyawan Swasta", "Explore Bahasa Go"},
	{2, "Afif Wanda Julio", "Cilegon", "Karyawan", "Mempelajari coding lebih lanjut"},
	{3, "Muhammad Nurul Afif Maliki", "Bandung", "Mahasiswa", "Banyak lowongan di LinkedIn"},
	{4, "Amanda Jehan", "Tangerang Selatan", "Fullstack Developer", "Banyak dicari perusahaan"},
	{5, "Gading Condro Prakoso", "Surakarta", "Freelance Developer", "Untuk project freelance selanjutnya"},
}

func main() {
	// Get the command-line arguments
	args := os.Args[1:]

	// Print usage message if no argument is provided
	if len(args) == 0 {
		fmt.Println("Gunakan go run main.go [Nama]")
		return
	}

	nama := args[0]

	// Attempt to convert the first argument to an integer (userId)
	userId, err := strconv.Atoi(args[0])

	// Check if the conversion fails or if userId is 0 (indicating it's not an integer)
	if err != nil {
		if userId == 0 {
			findFriendByName(nama)
		} else {
			fmt.Println("UserId harus berupa bilangan bulat")
		}
		return
	}

	found := false

	// Iterate over the friends slice to find the friend with the given userId
	for _, friend := range friends {
		if friend.Id == userId {
			found = true
			displayFriendData(friend)
			break
		}
	}

	// Print a message if the friend with the given userId is not found
	if !found {
		fmt.Println("Data dengan Id", userId, "tidak tersedia")
	}
}

// findFriendByName searches for a friend by their name and displays their data if found
func findFriendByName(nama string) {
	found := false
	for _, friend := range friends {
		if friend.Nama == nama {
			found = true
			displayFriendData(friend)
			break
		}
	}

	// Print a message if the friend with the given name is not found
	if !found {
		fmt.Println("Data dengan nama", nama, "tidak tersedia")
	}
}

// displayFriendData displays the details of a friend
func displayFriendData(friend Friend) {
	fmt.Println("Id:", friend.Id)
	fmt.Println("Nama:", friend.Nama)
	fmt.Println("Alamat:", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan memilih kelas golang:", friend.Alasan)
}

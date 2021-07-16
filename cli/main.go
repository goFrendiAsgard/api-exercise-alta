package main

import "fmt"

type User struct {
	Name    string
	Address string
}

func main() {
	var users []User = make([]User, 0)
	keluar := false
	for !keluar {
		fmt.Println("Masukkan pilihan: lihat/tambah/keluar")
		pilihan := ""
		fmt.Scanln(&pilihan)
		switch pilihan {
		case "lihat":
			for _, user := range users {
				fmt.Println("Name", user.Name)
				fmt.Println("Address", user.Address)
			}
		case "tambah":
			newUser := User{}
			fmt.Println("Masukkan nama")
			fmt.Scanln(&newUser.Name)
			fmt.Println("Masukkan alamat")
			fmt.Scanln(&newUser.Address)
			users = append(users, newUser)
		case "keluar":
			keluar = true
		}
	}
}

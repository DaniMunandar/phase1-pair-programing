package main

import (
	"fmt"
	"phas1-pair-programing/config"
	"phas1-pair-programing/entity"
	"phas1-pair-programing/handler"
)

func main() {
	config.ConnectDB()

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Println("Registrasi")
			fmt.Print("Masukan username:")
			fmt.Scanln(&username)
			fmt.Print("Masukan password:")
			fmt.Scanln(&password)

			customer := entity.Customer{
				Username: username,
				Password: password,
			}

			err := handler.Register(customer)
			if err != nil {
				fmt.Println("Kesalahan saat registrasi:", err)
			} else {
				fmt.Println("Berhasil Registrasi")
			}
		case 2:
			var username, password string
			fmt.Println("Login")
			fmt.Print("Masukan username:")
			fmt.Scanln(&username)
			fmt.Print("Masukan password:")
			fmt.Scanln(&password)

			isAuthenticated, err := handler.Login(username, password)
			if err != nil {
				fmt.Println("Kesalahan saat login:", err)
			} else if isAuthenticated {
				fmt.Println("Login berhasil! Selamat datang,", username)
			} else {
				fmt.Println("Login gagal")
			}
		case 3:
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

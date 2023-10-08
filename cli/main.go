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
		fmt.Println("***** Welcome to Library System *****")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")
		fmt.Println("-------------------------------------")

		var choice int
		fmt.Print("Choice menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username, password string

			fmt.Printf("\n*** Register Page ***\n")
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			customer := entity.Customer{
				Username: username,
				Password: password,
			}

			err := handler.Register(customer)
			if err != nil {
				fmt.Println("Kesalahan saat registrasi: ", err)
			} else {
				fmt.Printf("\n*** Berhasil Registrasi ***\n\n")
			}

		case 2:
			var username, password string
			fmt.Printf("\n*** Login Page ***\n")
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			isAdmin, err := handler.Login(username, password)
			if err != nil {
				fmt.Printf("\n*** Failed to login. Please try again. [Error] : %v\n", err)
			}

			if isAdmin {
				AdminPage(username)
			} else {
				UserPage(username)
			}

		case 0:
			fmt.Println("Keluar dari program")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

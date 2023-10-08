package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Println("3. Keluar")
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

			isAuthenticated, err := handler.Login(username, password)
			if err != nil {
				fmt.Println("Kesalahan saat login: ", err)
			} else if isAuthenticated {
				UserPage(username)
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

func UserPage(username string) {
	for {
		fmt.Printf("\n***** Welcome to User Page, %s *****\n", username)
		fmt.Println("1. Loan")
		fmt.Println("2. List Book")
		fmt.Println("3. Logout")
		fmt.Println("-----------------------------------------")

		var choice int
		fmt.Print("Choice menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Input Scanner
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println()
			fmt.Println("***** Loan System *****")

			fmt.Print("Book Title : ")
			scanner.Scan()
			title := scanner.Text()

			var duration int
			fmt.Print("Duration (days) : ")
			fmt.Scanln(&duration)

			msg, err := handler.Loan(title, username, duration)
			if err != nil {
				fmt.Printf("\n %s \n", msg)
			} else {
				fmt.Printf("\n %s \n", msg)
			}

		case 2:
			books, err := handler.GetBook()
			if err != nil {
				log.Fatal("Server Error : ", err)
			} else {
				fmt.Printf("\n\n***** List Book *****\n")
				fmt.Println("TITLE | AUTHOR | PUBLICATION DATE")
				fmt.Println("-------------------------------------------------")
				for _, book := range books {
					fmt.Printf("%s | %s | %s\n", book.Title, book.Author, book.PublicationDate)
				}
			}

		case 3:
			fmt.Printf("\n*** Logout Success ***\n\n")
			return

		default:
			fmt.Println("Invalid Input")
		}
	}
}

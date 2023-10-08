package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "phas1-pair-programing/entity"
	"phas1-pair-programing/handler"
	// "strconv"
	// "time"
)

func UserPage(username string) {
	for {
		fmt.Printf("\n***** Welcome to User Page, %s *****\n", username)
		fmt.Println("1. Loan")
		fmt.Println("2. List Book")
		fmt.Println("0. Logout")
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
				fmt.Println("| TITLE 		| AUTHOR 		| PUBLICATION DATE |")
				fmt.Println("--------------------------------------------------------------------")
				for _, book := range books {
					fmt.Printf("| %s 		| %s 		| %s |\n", book.Title, book.Author, book.PublicationDate)
				}
			}

		case 0:
			fmt.Printf("\n*** Logout Success ***\n\n")
			return

		default:
			fmt.Println("Invalid Input")
		}
	}
}

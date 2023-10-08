package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phas1-pair-programing/config"
	"phas1-pair-programing/entity"
	"phas1-pair-programing/handler"
	"strconv"
	"time"
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
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Add Book")
		fmt.Println("6. Logout")
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

		case 3:
			// Implementasi Update Book
			UpdateBook()
		case 4:
			// Implementasi Delete Book
			DeleteBook()
		case 5:
			// Implementasi Delete Book
			AddBook()
		case 6:
			fmt.Printf("\n*** Logout Success ***\n\n")
			return
		default:
			fmt.Println("Invalid Input")
		}
	}
}

func UpdateBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("***** Update Book *****")
	fmt.Print("Book ID: ")
	scanner.Scan()
	bookIDStr := scanner.Text()
	bookID, _ := strconv.Atoi(bookIDStr)

	// Mendapatkan informasi buku yang akan diperbarui
	book, err := handler.GetBookByID(bookID)
	if err != nil {
		fmt.Println("Kesalahan saat mengambil informasi buku: ", err)
		return
	}

	fmt.Printf("Current Title: %s\n", book.Title)
	fmt.Printf("Current Author: %s\n", book.Author)
	fmt.Printf("Current Publication Date: %s\n", book.PublicationDate)

	fmt.Print("New Title: ")
	scanner.Scan()
	newTitle := scanner.Text()

	fmt.Print("New Author: ")
	scanner.Scan()
	newAuthorName := scanner.Text()

	// Memperbarui objek buku dengan data baru
	book.Title = newTitle
	book.Author = newAuthorName // Menggunakan nama penulis yang baru
	fmt.Print("New Publication Date: ")
	scanner.Scan()
	newPublicationDate := scanner.Text()

	book.PublicationDate = newPublicationDate

	// Memperbarui data dalam database dengan fungsi UpdateBook
	err = handler.UpdateBook(book)
	if err != nil {
		fmt.Println("Kesalahan saat memperbarui buku: ", err)
	} else {
		fmt.Printf("Buku dengan ID %d telah diperbarui.\n", bookID)
	}
}

func DeleteBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("***** Delete Book *****")
	fmt.Print("Book ID: ")
	scanner.Scan()
	bookIDStr := scanner.Text()
	bookID, _ := strconv.Atoi(bookIDStr)

	// Memanggil fungsi DeleteBook dari handler
	err := handler.DeleteBook(bookID)
	if err != nil {
		fmt.Println("Kesalahan saat menghapus buku: ", err)
	} else {
		fmt.Printf("Buku dengan ID %d telah dihapus.\n", bookID)
	}
}

func AddBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("***** Tambah Buku *****")

	fmt.Print("Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Author ID: ") // Meminta pengguna untuk memasukkan ID penulis yang sesuai
	scanner.Scan()
	authorIDStr := scanner.Text()
	authorID, _ := strconv.Atoi(authorIDStr)

	fmt.Print("Publication Date (YYYY-MM-DD): ")
	scanner.Scan()
	publicationDateStr := scanner.Text()

	// Menguji validitas format tanggal
	parsedDate, err := time.Parse("2006-01-02", publicationDateStr)
	if err != nil {
		fmt.Println("Kesalahan format tanggal. Gunakan format YYYY-MM-DD.")
		return
	}
	// Menggunakan tanggal yang sudah diuji validitasnya
	newPublicationDate := parsedDate.Format("2006-01-02")

	// Memasukkan data buku ke dalam database dengan mengatur author_id sesuai dengan yang dimasukkan pengguna
	newBook := entity.Book{
		Title:           title,
		AuthorID:        authorID,
		PublicationDate: newPublicationDate,
	}

	// Memanggil fungsi AddBook dari handler
	err = handler.AddBook(newBook)
	if err != nil {
		fmt.Println("Kesalahan saat menambah buku: ", err)
	} else {
		fmt.Println("Buku berhasil ditambahkan.")
	}
}

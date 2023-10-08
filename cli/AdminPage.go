package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phas1-pair-programing/entity"
	"phas1-pair-programing/handler"
	"strconv"
	"time"
)

func AdminPage(username string) {
	for {
		fmt.Printf("\n***** Welcome to Admin Page, %s *****\n", username)
		fmt.Println("1. List Book")
		fmt.Println("2. Update Book")
		fmt.Println("3. Delete Book")
		fmt.Println("4. Add Book")
		fmt.Println("0. Logout")
		fmt.Println("-----------------------------------------")

		var choice int
		fmt.Print("Choice menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			books, err := handler.GetBook()
			if err != nil {
				log.Fatal("Server Error : ", err)
			} else {
				fmt.Printf("\n\n***** List Book *****\n")
				fmt.Println("| TITLE			| AUTHOR		| PUBLICATION DATE	|")
				fmt.Println("--------------------------------------------------------------------------")
				for _, book := range books {
					fmt.Printf("| %s		| %s		| %s		|\n", book.Title, book.Author, book.PublicationDate)
				}
			}

		case 2:
			UpdateBook()
		case 3:
			DeleteBook()
		case 4:
			AddBook()
		case 0:
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

	fmt.Print("New Author ID : ")
	scanner.Scan()
	newAuthorName := scanner.Text()

	// Memperbarui objek buku dengan data baru
	book.Title = newTitle
	authorId, err := strconv.Atoi(newAuthorName)
	if err != nil {
		log.Fatal("Invalid Input")
	}
	book.AuthorID = authorId // ID Author
	fmt.Print("New Publication Date: ")
	scanner.Scan()
	newPublicationDate := scanner.Text()

	book.PublicationDate = newPublicationDate

	// Memperbarui data dalam database dengan fungsi UpdateBook
	err = handler.UpdateBook(&book)
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

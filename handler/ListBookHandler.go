package handler

import (
	"context"
	"phas1-pair-programing/config"
	"phas1-pair-programing/entity"
	"time"
)

func GetBook() ([]entity.Book, error) {
	var books []entity.Book

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := config.DB.QueryContext(ctx, "SELECT books.title AS 'Title', author.name AS 'Author', books.publication_date AS 'Publication' FROM books JOIN author ON books.author_id = books.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b entity.Book

		if err := rows.Scan(&b.Title, &b.Author, &b.PublicationDate); err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

func CreateBook(book entity.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DB.ExecContext(ctx, "INSERT INTO books (title, author_id, publication_date) VALUES (?, ?, ?)", book.Title, book.Author, book.PublicationDate)
	if err != nil {
		return err
	}

	return nil
}

func GetAuthorIDByName(authorName string) (int, error) {
	var authorID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := config.DB.QueryRowContext(ctx, "SELECT id FROM author WHERE name = ?", authorName).Scan(&authorID)
	if err != nil {
		return 0, err
	}

	return authorID, nil
}

func GetBookByID(bookID int) (entity.Book, error) {
	var book entity.Book

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := config.DB.QueryRowContext(ctx, "SELECT title, author_id, publication_date FROM books WHERE id = ?", bookID).Scan(&book.Title, &book.Author, &book.PublicationDate)
	if err != nil {
		return book, err
	}

	return book, nil
}

func UpdateBook(book entity.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gunakan query SQL UPDATE untuk memperbarui buku dalam database.
	_, err := config.DB.ExecContext(ctx, "UPDATE books SET title=?, author_id=?, publication_date=? WHERE id=?", book.Title, book.AuthorID, book.PublicationDate, book.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(bookID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Mulai transaksi
	tx, err := config.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback transaksi jika terjadi kesalahan

	// Menghapus peminjaman terkait buku
	_, err = tx.ExecContext(ctx, "DELETE FROM loan WHERE book_id = ?", bookID)
	if err != nil {
		return err
	}

	// Menghapus buku
	_, err = tx.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		return err
	}

	// Commit transaksi jika berhasil
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func AddBook(book entity.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tambahkan buku ke database
	_, err := config.DB.ExecContext(ctx, "INSERT INTO books (title, author_id, publication_date) VALUES (?, ?, ?)", book.Title, book.AuthorID, book.PublicationDate)
	if err != nil {
		return err
	}

	return nil
}

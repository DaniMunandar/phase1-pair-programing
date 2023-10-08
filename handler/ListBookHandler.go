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

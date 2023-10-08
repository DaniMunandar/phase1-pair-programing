package handler

import (
	"context"

	"phas1-pair-programing/config"
	"time"
)

func Loan(title, username string, duration int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book_id int

	// Check if book exist
	err := config.DB.QueryRowContext(ctx, "SELECT id FROM books WHERE title = ?", title).Scan(&book_id)
	if err != nil {
		return "*** Book Title Not Found ***", err
	}

	// Get Customer Data
	var customer_id int
	err = config.DB.QueryRowContext(ctx, "SELECT id FROM customers WHERE username = ?", username).Scan(&customer_id)
	if err != nil {
		return "*** Customer Data Not Found ***", err
	}

	_, err = config.DB.ExecContext(ctx, "INSERT INTO loan (customers_id, book_id, loan_date, return_date) VALUES (?, ?, CURDATE(), CURDATE() + ?)", customer_id, book_id, duration)
	if err != nil {
		return "*** Failed iserting loan data ***", err
	}

	return "*** Loan Success ***", err
}

package handler

import "database/sql"

func Register(db *sql.DB) error {
	db.Exec("INSERT INTO customers (username, password)")

}

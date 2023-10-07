// package config

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func GetDB() (*sql.DB, error) {
// 	connStr := "root:@tcp(localhost:3306)/buku"
// 	db, err := sql.Open("mysql", connStr)
// 	return db, err
// }

// config/config.go

package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Ganti dengan driver yang sesuai
)

var DB *sql.DB

// Inisialisasi DB
func ConnectDB() {
	// Gantilah dengan pengaturan koneksi ke database Anda
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/buku")
	if err != nil {
		log.Fatal(err)
	}

	// Tes koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

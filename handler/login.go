package handler

import (
	"phas1-pair-programing/config"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (bool, error) {
	var hashedPassword string
	var isAdmin bool

	// Cari pelanggan berdasarkan nama pengguna
	err := config.DB.QueryRow("SELECT password, isAdmin FROM customers WHERE username = ?", username).Scan(&hashedPassword, &isAdmin)
	if err != nil {
		return isAdmin, err
	}

	// Verifikasi kata sandi
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return isAdmin, err // Kata sandi tidak cocok
	}

	return isAdmin, nil // Otentikasi berhasil
}

package handler

import (
	"phas1-pair-programing/config"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (bool, error) {
	var hashedPassword string

	// Cari pelanggan berdasarkan nama pengguna
	err := config.DB.QueryRow("SELECT password FROM customers WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	// Verifikasi kata sandi
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil // Kata sandi tidak cocok
	}

	return true, nil // Otentikasi berhasil
}

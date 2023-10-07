package handler

import (
	"phas1-pair-programing/config"
	"phas1-pair-programing/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(customer entity.Customer) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO customers (username, password) VALUES (?,?)", customer.Username, hashedPassword)
	return err
}

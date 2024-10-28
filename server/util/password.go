package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bcyptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password")
	}
	return string(bcyptPassword), nil

}

func CheckPassword(password string,hashedpassword string)error{
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword),[]byte(password))

}
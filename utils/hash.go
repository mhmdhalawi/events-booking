package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bcryptPassword)
}

func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

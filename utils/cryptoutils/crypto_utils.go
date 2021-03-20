package cryptoutils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Encrypt(in string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(in), 13)
	if err != nil {
		log.Fatal(err)
	}
	return string(password)
}

func VerifyPassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePass := []byte(plainPwd)

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePass); err != nil {
		return false
	}
	return true
}

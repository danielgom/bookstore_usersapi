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

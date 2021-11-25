package hashing

import "golang.org/x/crypto/bcrypt"

func Generate(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(res), err
}

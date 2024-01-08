package utils

import (
	"filip.filipovic/polling-app/logging"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword() Uses the BCrypt algorithm to return the hash value of the given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

//MustHashPassword() Same as HashPassword() but panics if there is an error
func MustHashPassword(password string) string {
	hash, err := HashPassword(password)
	if err != nil {
		logging.Panic(err)
	}
	return hash
}

// CheckPasswordHash() returns true if the hash matches the password, or false if it doesn't
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

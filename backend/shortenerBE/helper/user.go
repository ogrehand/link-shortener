package helper

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func CheckHash(salt string, password string, hash string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	if result == nil {
		return true
	} else {
		return false
	}
}

func EncryptPassword(salt string, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(bytes), err
}

func GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

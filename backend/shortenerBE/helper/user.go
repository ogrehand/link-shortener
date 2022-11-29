package helper

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(salt string, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(bytes), err
}

func GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CheckToken(id string) {

}

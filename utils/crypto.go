package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const cost int = 16

type cryptoUtils struct {
	Cost   int
	Hash   func(password string) (string, error)
	Equals func(hashedPassword string, password string) bool
}

func Crypto() cryptoUtils {
	return cryptoUtils{
		Cost:   cost,
		Hash:   hash,
		Equals: equals,
	}
}

func hash(password string) (string, error) {
	newPassword := []byte(password)
	hash, exception := bcrypt.GenerateFromPassword(newPassword, cost)

	if exception != nil {
		return "", fmt.Errorf("utils.crypto: %s", exception.Error())
	}

	return string(hash), exception
}

func equals(hashedPassword string, password string) bool {
	newHashedPassword := []byte(hashedPassword)
	newPassword := []byte(password)

	exception := bcrypt.CompareHashAndPassword(
		newHashedPassword, newPassword)

	return exception == nil
}

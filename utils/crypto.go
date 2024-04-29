package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const cost int = 16

type CryptoUtils struct {
	Cost   int
	Hash   func(password string) (string, error)
	Equals func(hashedPassword string, password string) bool
	Serial func(input string, length int) string
}

func Crypto() CryptoUtils {
	return CryptoUtils{
		Cost:   cost,
		Hash:   hash,
		Equals: equals,
		Serial: serial,
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

func serial(input string, length int) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	base64Str := base64.StdEncoding.EncodeToString(hashBytes)

	urlSafeStr := strings.ReplaceAll(base64Str, "+", "-")
	urlSafeStr = strings.ReplaceAll(urlSafeStr, "/", "_")

	urlSafeStr = strings.TrimRight(urlSafeStr, "=")

	return urlSafeStr[:length]
}

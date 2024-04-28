package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
)

type UserUtil struct {
	GetLoginToken func(hash string) string
	GenerateHash  func(input string) string
}

func User() UserUtil {
	return UserUtil{
		GetLoginToken: getLoginToken,
		GenerateHash:  generateHash,
	}
}

func getLoginToken(hash string) string {
	return hash + os.Getenv("TOKEN_SECRET")
}

func generateHash(input string) string {
	length := 16

	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	base64Str := base64.StdEncoding.EncodeToString(hashBytes)

	urlSafeStr := strings.ReplaceAll(base64Str, "+", "-")
	urlSafeStr = strings.ReplaceAll(urlSafeStr, "/", "_")

	urlSafeStr = strings.TrimRight(urlSafeStr, "=")

	return urlSafeStr[:length]
}

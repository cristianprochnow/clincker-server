package utils

import (
	"os"
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
	return Crypto().Serial(input, 16)
}

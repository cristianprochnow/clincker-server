package utils

import (
	"net/mail"
	"os"
)

type UserUtil struct {
	GetLoginToken func(hash string) string
	GenerateHash  func(input string) string
	IsValidEmail  func(email string) bool
	IsTooLongPass func(password string) bool
	PassSize      int
	NoPassValue   string
}

func User() UserUtil {
	return UserUtil{
		GetLoginToken: getLoginToken,
		GenerateHash:  generateHash,
		IsValidEmail:  isValidEmail,
		IsTooLongPass: isTooLongPassword,
		PassSize:      PASS_SIZE,
		NoPassValue:   NO_PASS_VALUE,
	}
}

const PASS_SIZE int = 16
const NO_PASS_VALUE string = "NO-PASS"

func getLoginToken(hash string) string {
	return hash + os.Getenv("TOKEN_SECRET")
}

func generateHash(input string) string {
	return Crypto().Serial(input, 16)
}

func isValidEmail(email string) bool {
	if len(email) == 0 {
		return false
	}

	emailAddress, exception := mail.ParseAddress(email)

	return exception == nil && emailAddress.Address == email
}

func isTooLongPassword(password string) bool {
	return len(password) > PASS_SIZE
}

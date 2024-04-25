package utils

import "os"

type UserUtil struct {
	GetLoginToken func(email string, user string) string
}

func User() UserUtil {
	return UserUtil{
		GetLoginToken: getLoginToken,
	}
}

func getLoginToken(email string, user string) string {
	return email + user + os.Getenv("TOKEN_SECRET")
}

package utils

import (
	"log"
	"strings"
)

type LogUtil struct {
	Exception func(exception error)
}

func Log() LogUtil {
	return LogUtil{
		Exception: exception,
	}
}

func IsNoRowsError(errorText string) bool {
	return strings.Contains(errorText, "no rows in result set")
}

func exception(exception error) {
	if exception != nil {
		log.Fatal(exception)
	}
}

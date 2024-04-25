package utils

import (
	"log"
	"strings"
)

type LogUtil struct {
	Exception     func(exception error)
	IsNoRowsError func(errorText string) bool
}

func Log() LogUtil {
	return LogUtil{
		Exception:     exception,
		IsNoRowsError: isNoRowsError,
	}
}

func isNoRowsError(errorText string) bool {
	return strings.Contains(errorText, "no rows in result set")
}

func exception(exception error) {
	if exception != nil {
		log.Fatal(exception)
	}
}

package utils

import "log"

type LogUtil struct {
	Exception func(exception error)
}

func Log() LogUtil {
	return LogUtil{
		Exception: exception,
	}
}

func exception(exception error) {
	if exception != nil {
		log.Fatal(exception)
	}
}

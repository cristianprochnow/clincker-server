package utils

import "log"

func Error(exception error) {
	if exception != nil {
		log.Fatal(exception)
	}
}

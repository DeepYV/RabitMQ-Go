package util

import "log"

func Error(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

package helper

import (
	"log"
)

func handlerError(err *error) {
	if err := recover(); err != nil {
		log.Println("error recover:", err)
	}
}

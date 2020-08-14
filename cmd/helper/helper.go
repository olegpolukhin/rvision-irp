package helper

import (
	log "github.com/sirupsen/logrus"
)

func handlerError(err *error) {
	if err := recover(); err != nil {
		log.Error("error recover:", err)
	}
}

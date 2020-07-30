package datasource

import (
	"log"
	udb "upper.io/db.v3"
)

func PanicOnError(err error) error {
	if err != nil && err != udb.ErrNoMoreRows {
		log.Println("Error: ", err)
		panic(err)
	}
	if err == udb.ErrNoMoreRows {
		return nil
	}
	return err
}

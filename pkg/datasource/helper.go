package datasource

import (
	"log"
	udb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

func executeInsertAndGetNewID(inserter sqlbuilder.Inserter) (newID string, err error) {
	iter := inserter.Returning("id").Iterator()
	if err = iter.ScanOne(&newID); err != nil {
		return
	}
	return
}

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

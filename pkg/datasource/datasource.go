package datasource

import (
	"context"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

type Postgres struct {
	pgURI postgresql.ConnectionURL
	db    sqlbuilder.Database
	tx    sqlbuilder.Tx
}

func NewPgDatasource(uri postgresql.ConnectionURL) (pg *Postgres) {
	pg = &Postgres{
		pgURI: uri,
		db:    nil,
	}

	return pg
}

func (pg *Postgres) Connect() (err error) {
	pg.db, err = postgresql.Open(pg.pgURI)
	if err != nil {
		return err
	}

	if err = pg.db.Ping(); err != nil {
		return err
	}

	tx, err := pg.db.NewTx(context.Background())
	if err != nil {
		return err
	}
	pg.tx = tx

	return
}

func (pg *Postgres) Close(err *error) interface{} {
	// Проверяем на наличие ошибок, если есть, откатываемся
	if *err != nil {
		// Если в ошибке что-то уже есть, то откатываемся
		pg.tx.Rollback()
		return nil
	}
	if p := recover(); p != nil {
		// Читаем информацию по ошибке для отправки на фронт
		*err = p.(error)
		// Откатываем базу
		pg.tx.Rollback()
		return p
	} else {
		pg.tx.Commit()
	}
	// Закрываем соединение с базой
	return pg.db.Close()
}

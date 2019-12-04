package repository

import (
	"database/sql"
	"runtime"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type DataSource struct {
	*reform.DB
}

func (d DataSource) Close() {
	err := d.DBInterface().(*sql.DB).Close()
	if err != nil {
		log.Error().Err(err).Msg("DataSource")
	}
}

func NewDataSource(dbConnectionString string) (ds DataSource, err error) {
	config, err := pgx.ParseConfig(dbConnectionString)
	if err != nil {
		return
	}
	db := stdlib.OpenDB(*config)
	db.SetMaxOpenConns(runtime.NumCPU() * 4)
	if err = db.Ping(); err != nil {
		return
	}
	ds.DB = reform.NewDB(db, postgresql.Dialect, nil)
	return
}

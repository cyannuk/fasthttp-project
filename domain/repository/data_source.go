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

type dataSource struct {
	*reform.DB
}

func (d dataSource) CloseDataSource() {
	err := d.DBInterface().(*sql.DB).Close()
	if err != nil {
		log.Error().Err(err)
	}
}

func NewDataSource(dbConnectionString string) (dataSource, error) {
	config, err := pgx.ParseConfig(dbConnectionString)
	if err != nil {
		return dataSource{}, err
	}
	db := stdlib.OpenDB(*config)
	db.SetMaxOpenConns(runtime.NumCPU() * 4)
	return dataSource{reform.NewDB(db, postgresql.Dialect, nil)}, nil
}

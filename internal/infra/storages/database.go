package storages

import (
	"gopaseto/internal/core/config"
	"gopaseto/internal/core/port/repository"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	*sqlx.DB
}

// GetDB implements repository.Database.
func (db database) GetDB() *sqlx.DB {
	return db.DB
}

func (db database) Close() error {
	return db.DB.Close()
}

func NewDB(conf config.DB) (repository.Database, error) {
	db, err := newDatabase(conf)

	if err != nil {
		return nil, err
	}

	return &database{
		db,
	}, nil
}

func newDatabase(conf config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(conf.Driver, conf.Url)

	if err != nil {
		return nil, err
	}

	maxLifeTime, errLifeTime := strconv.Atoi(conf.MaxLifeTime)
	maxOpenConns, errOpenConns := strconv.Atoi(conf.MaxOpenConn)
	maxIdleConn, errMaxIdle := strconv.Atoi(conf.MaxIddleConn)

	if errLifeTime != nil && errOpenConns != nil && errMaxIdle != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(maxLifeTime))
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConn)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

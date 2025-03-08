package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

func NewConn(dsn string, maxOpenConn, maxIdleConn int, maxIdleTimeout string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// pooling configurations.
	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)

	duration, er := time.ParseDuration(maxIdleTimeout)
	if er != nil {
		return nil, errors.New("unable to parse time duration")
	}
	db.SetConnMaxLifetime(duration)

	// ping database.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // receives a cancellation signal.
	defer cancel()

	if e := db.PingContext(ctx); e != nil {
		return nil, e
	}

	log.Println("Connected to database successfully.")

	return db, nil

}

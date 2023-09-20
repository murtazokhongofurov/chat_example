package db

import (
	"database/sql"
	"fmt"

	"github.com/kafka_example/chat_service/config"
	_ "github.com/lib/pq"
)

func ConnectToDb(cfg config.Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	conn, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func ConnectToDbSuiteTest(cfg config.Config) (*sql.DB, func()) {
	psqlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	conn, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, func() {}
	}
	cleanUp := func() {

	}
	return conn, cleanUp
}

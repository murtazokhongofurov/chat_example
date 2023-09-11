package postgres

import (
	"database/sql"

	"github.com/kafka_example/chat_service/storage/models"
)

type storagePg struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storagePg {
	return &storagePg{
		db: db,
	}
}

func (p *storagePg) SavePerson(person models.Message) error {
	query := "INSERT INTO person(name, age) VALUES($1, $2)"
	_, err := p.db.Exec(query, person.Name, person.Age)
	return err
}

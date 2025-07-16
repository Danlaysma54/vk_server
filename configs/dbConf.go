package configs

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type Storage struct {
	Db *sql.DB
}

func runMigrations(dbSourceURL string, dbURL string) {
	m, err := migrate.New(dbSourceURL, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
	}
	log.Println("Migrations applied successfully")
}

func NewStorage() (*Storage, error) {
	LoadEnv()
	runMigrations(getSourceURL(), getDbURL())
	const op = "NewStorage"
	db, err := sql.Open("postgres", getConnString())
	if err != nil {
		return nil, fmt.Errorf("%s - %s", op, err)
	}

	return &Storage{Db: db}, nil
}

func CloseDB(storage *Storage) {
	storage.Db.Close()
}

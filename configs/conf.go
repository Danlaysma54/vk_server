package configs

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"time"
)

type JWT struct {
	Secret      string
	TokenExpiry time.Duration
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

func NewStorage() (*sql.DB, error) {
	LoadEnv()
	runMigrations(getSourceURL(), getDbURL())
	const op = "NewStorage"
	db, err := sql.Open("postgres", getConnString())
	if err != nil {
		return nil, fmt.Errorf("%s - %s", op, err)
	}

	return db, nil
}
func JwtConfig() JWT {
	j := JWT{}
	j.Secret = getJWT_SECRET()
	j.TokenExpiry = time.Hour * 24
	return j
}

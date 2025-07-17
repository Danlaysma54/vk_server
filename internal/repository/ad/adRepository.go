package ad

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type RepoAd struct {
	db *sql.DB
}

func NewRepoAd(db *sql.DB) *RepoAd {
	return &RepoAd{
		db: db,
	}
}
func (s *RepoAd) SaveAd(
	adName string,
	description string,
	price int,
	authorId string) uuid.UUID {
	var res uuid.UUID
	stmt, err2 := s.db.Prepare(
		"INSERT INTO ad(ad_name,description,price,author_id) VALUES ($1,$2,$3,$4) returning id")
	if err2 != nil {
		panic(err2)
	}
	err := stmt.QueryRow(adName, description, price, authorId).Scan(&res)
	if err != nil {
		log.Fatal("Can't insert ad")
	}
	return res
}

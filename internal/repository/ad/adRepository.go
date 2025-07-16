package ad

import (
	"github.com/google/uuid"
	"log"
	"vk_server/configs"
)

type AddStorage struct {
	*configs.Storage
}

func NewAddStorage(storage *configs.Storage) *AddStorage {
	return &AddStorage{
		storage,
	}
}
func (s *AddStorage) SaveAd(
	adName string,
	description string,
	price int,
	authorId string) uuid.UUID {
	var res uuid.UUID
	stmt, err2 := s.Storage.Db.Prepare(
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

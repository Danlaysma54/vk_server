package ad

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"vk_server/internal/model"
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
	imageUrl string,
	price int,
	authorId string) uuid.UUID {
	var res uuid.UUID
	stmt, err2 := s.db.Prepare(
		"INSERT INTO ad(ad_name,description,image_url,price,author_id) VALUES ($1,$2,$3,$4,$5) returning id")
	if err2 != nil {
		println(err2)
	}
	err := stmt.QueryRow(adName, description, imageUrl, price, authorId).Scan(&res)
	if err != nil {
		log.Fatal("Can't insert ad")
	}
	return res
}
func (s *RepoAd) GetAd(adId uuid.UUID) *model.AdEntity {
	stmt, err := s.db.Prepare("SELECT id,ad_name,description,image_url,price,author_id FROM ad where  id=$1")
	if err != nil {
		println(err)
	}
	var adEnt model.AdEntity
	stmt.QueryRow(adId).Scan(&adEnt.AdId, &adEnt.Name, &adEnt.Description, &adEnt.ImageUrl, &adEnt.Price, &adEnt.AuthorId)
	return &adEnt
}

package ad

import (
	"database/sql"
	"fmt"
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
	stmt, err := s.db.Prepare("SELECT ad.id,ad.ad_name,ad.description,ad.image_url,ad.price,users.username  FROM ad inner join users on users.id=ad.author_id where  ad.id=$1")
	if err != nil {
		println(err)
	}
	var adEnt model.AdEntity
	stmt.QueryRow(adId).Scan(&adEnt.AdId, &adEnt.Name, &adEnt.Description, &adEnt.ImageUrl, &adEnt.Price, &adEnt.Username)
	return &adEnt
}
func (s *RepoAd) GetAllAds(
	DateSort string,
	PriceSort string,
	MinPrice int,
	MaxPrice int,
	Limit int,
	Offset int) []*model.AdEntity {
	stmt := fmt.Sprintf(`select ad.id,ad.ad_name,ad.description,ad.image_url,ad.price,users.username FROM ad inner join users on users.id=ad.author_id where ad.price >$1 and ad.price< $2 order by ad.created_at %s, ad.price %s  limit $3 offset $4`,
		DateSort, PriceSort)
	rows, err := s.db.Query(stmt, &MinPrice, &MaxPrice, &Limit, Limit*Offset)
	if err != nil {
		println(err)
	}
	defer rows.Close()
	var AllEnts = []*model.AdEntity{}
	for rows.Next() {
		var AdEnt model.AdEntity
		err := rows.Scan(&AdEnt.AdId, &AdEnt.Name, &AdEnt.Description, &AdEnt.ImageUrl, &AdEnt.Price, &AdEnt.Username)
		if err != nil {
			println(err)
		}
		AllEnts = append(AllEnts, &AdEnt)
	}
	return AllEnts
}
func (s *RepoAd) GetAllAdsForAuth(
	DateSort string,
	PriceSort string,
	MinPrice int,
	MaxPrice int,
	Limit int,
	Offset int,
	AuthorId string) []*model.AdEntityForAuth {
	stmt := fmt.Sprintf(`select ad.id,ad.ad_name,ad.description,ad.image_url,ad.price,users.username,ad.author_id FROM ad inner join users on users.id=ad.author_id where ad.price >$1 and ad.price< $2 order by ad.created_at %s, ad.price %s  limit $3 offset $4`,
		DateSort, PriceSort)
	rows, err := s.db.Query(stmt, &MinPrice, &MaxPrice, &Limit, Limit*Offset)
	if err != nil {
		println(err)
	}
	defer rows.Close()
	var AllEnts = []*model.AdEntityForAuth{}
	for rows.Next() {
		var AdEnt model.AdEntityForAuth
		var AuthorAdId string
		err := rows.Scan(&AdEnt.AdId, &AdEnt.Name, &AdEnt.Description, &AdEnt.ImageUrl, &AdEnt.Price, &AdEnt.Username, &AuthorAdId)
		if AuthorId == AuthorAdId {
			AdEnt.Mine = true
		} else {
			AdEnt.Mine = false
		}
		if err != nil {
			println(err)
		}
		AllEnts = append(AllEnts, &AdEnt)
	}
	return AllEnts
}

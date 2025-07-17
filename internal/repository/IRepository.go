package repository

import (
	"github.com/google/uuid"
	"vk_server/internal/model"
)

type IRepoAd interface {
	SaveAd(
		adName string,
		description string,
		imageUrl string,
		price int,
		authorId string) uuid.UUID
	GetAd(adId uuid.UUID) *model.AdEntity
}

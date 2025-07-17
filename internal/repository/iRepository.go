package repository

import "github.com/google/uuid"

type IRepoAd interface {
	SaveAd(
		adName string,
		description string,
		price int,
		authorId string) uuid.UUID
}

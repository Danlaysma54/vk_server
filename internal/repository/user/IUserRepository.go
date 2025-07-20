package user

import "vk_server/internal/model"

type IRepoUser interface {
	IsExist(username string) bool
	InsertUser(username string, password string)
	GetUserByUsername(username string) *model.User
}

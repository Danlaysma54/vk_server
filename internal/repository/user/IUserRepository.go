package user

type IRepoUser interface {
	GetUser(username string) bool
	InsertUser(username string, password string)
}

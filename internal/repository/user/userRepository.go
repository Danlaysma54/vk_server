package user

import (
	"database/sql"
	"vk_server/internal/model"
)

type RepoUser struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *RepoUser {
	return &RepoUser{
		db: db,
	}
}
func (r *RepoUser) IsExist(username string) bool {
	var exists bool
	stmt, err := r.db.Prepare("SELECT EXISTS(SELECT username FROM users WHERE username = $1)")
	if err != nil {
		panic(err)
	}
	stmt.QueryRow(username).Scan(&exists)
	return exists
}

func (r *RepoUser) InsertUser(username string, password string) {
	stmt, err := r.db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(username, password)
}
func (r *RepoUser) GetUserByUsername(username string) *model.User {
	var user model.User
	stmt, err := r.db.Prepare("SELECT id,username, password FROM users WHERE username = $1")
	if err != nil {
		panic(err)
	}
	stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Password)
	return &user
}

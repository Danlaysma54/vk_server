package user

import "database/sql"

type RepoUser struct {
	db *sql.DB
}

func newUserRepo(db *sql.DB) *RepoUser {
	return &RepoUser{
		db: db,
	}
}
func (r *RepoUser) GetUser(username string) {
	var exists bool
	stmt, err := r.db.Prepare("SELECT username FROM users WHERE username = $1")
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

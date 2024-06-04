package models

type User struct {
	ID           int    `json:"id" bun:",pk,autoincrement"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

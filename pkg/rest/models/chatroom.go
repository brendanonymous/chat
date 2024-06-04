package models

type Chatroom struct {
	ID   int    `json:"id" bun:",pk,autoincrement"`
	Name string `json:"name"`
}

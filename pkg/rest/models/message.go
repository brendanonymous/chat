package models

import "time"

type Message struct {
	ID         int       `json:"id" bun:",pk,autoincrement"`
	UserID     int       `json:"user_id"`
	ChatroomID int       `json:"chatroom_id"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
}

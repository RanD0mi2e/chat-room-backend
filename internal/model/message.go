package model

type Message struct {
	ID          int    `db:"id"`
	Content     string `db:"content"`
	Author      User
	Attachments []Attachment
	Channel     Channel
	DeleteFlag  bool `db:"delete_flag"`
}

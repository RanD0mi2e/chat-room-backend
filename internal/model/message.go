package model

type Message struct {
	ID          int          `db:"id"`
	Content     string       `db:"content"`
	Author      User         `db:"-"`
	Attachments []Attachment `db:"-"`
	Channel     Channel      `db:"-"`
	DeleteFlag  bool         `db:"delete_flag"`
}

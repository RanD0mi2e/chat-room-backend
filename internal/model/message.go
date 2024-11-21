package model

type Message struct {
	ID          int          `db:"id"`
	ChannelId   string       `db:"channel_id"`
	Content     string       `db:"content"`
	Author      User         `db:"-"`
	Attachments []Attachment `db:"-"`
	DeleteFlag  bool         `db:"delete_flag"`
}

package model

type Channel struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Server Server
}

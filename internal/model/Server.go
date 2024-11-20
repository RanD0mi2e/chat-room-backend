package model

type Server struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Users    []User    `db:"-"`
	Channels []Channel `db:"-"`
}

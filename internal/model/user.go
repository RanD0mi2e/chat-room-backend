package model

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Accout    string `db:"accout"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	AvatarUrl string `db:"avatar_url"`
	Server    Server
}

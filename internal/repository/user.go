package repository

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

// 小写struct隐藏repo内部敏感变量
type userRepo struct {
	*Repository
}

// 通过接口向外暴露操作user表的api
type UserRepo interface {
	GetUserById(id int) (*User, error)
}

func NewUserRepo(repo *Repository) UserRepo {
	return &userRepo{
		Repository: repo,
	}
}

// 根据id查询用户信息
func (r *userRepo) GetUserById(id int) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id, name, email FROM users WHERE id = ?", id)
	return &user, err
}

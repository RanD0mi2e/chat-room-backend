package repository

import "chat-app-backend-service/internal/model"

// 小写struct隐藏repo内部敏感变量
type userRepo struct {
	*Repository
}

// 通过接口向外暴露操作user表的api
type UserRepo interface {
	GetUserById(id int) (*model.User, error)
}

func NewUserRepo(repo *Repository) UserRepo {
	return &userRepo{
		Repository: repo,
	}
}

// 根据id查询用户信息
func (r *userRepo) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT id, name, email FROM users WHERE id = ?", id)
	return &user, err
}

package infra

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type UserRepository struct {
	userList []domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

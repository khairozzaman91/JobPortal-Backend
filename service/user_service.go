package service

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type UserService interface {
	Login(email, password string) (*domain.User, error)
	Store(user domain.User) (domain.User, error)
	List() ([]domain.User, error)
	Get(id int) (*domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
	DeleteAll() error
}
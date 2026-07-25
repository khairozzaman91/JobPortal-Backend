package repository

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type UserRepository interface {
	Store(user domain.User) (domain.User, error)
	List() ([]domain.User, error)
	Get(id int) (*domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
	DeleteAll() error
}

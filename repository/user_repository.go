package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

type UserRepository interface {
	GetByEmail(email string) (*domain.User, error)
	Store(user domain.User) (domain.User, error)
	List() ([]domain.User, error)
	Get(id int) (*domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
	DeleteAll() error
}

type UserRepositoryImpl struct {
	db       *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
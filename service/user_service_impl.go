package service

import (
	"errors"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/repository"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) Login(email, password string) (*domain.User, error) {

	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}

	return nil, errors.New("invalid email or password")
}

func (s *UserServiceImpl) Store(user domain.User) (domain.User, error) {
	return s.repo.Store(user)
}

func (s *UserServiceImpl) List() ([]domain.User, error) {
	return s.repo.List()
}

func (s *UserServiceImpl) Get(id int) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *UserServiceImpl) Update(user domain.User) (domain.User, error) {
	return s.repo.Update(user)
}

func (s *UserServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserServiceImpl) DeleteAll() error {
	return s.repo.DeleteAll()
}
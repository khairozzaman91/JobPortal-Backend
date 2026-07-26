package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

type JobRepository interface {
	Store(job domain.Job) (domain.Job, error)
	List() ([]domain.Job, error)
	Get(id int) (*domain.Job, error)
	Update(job domain.Job) (domain.Job, error)
	Delete(id uint) error
	DeleteAll() error
}

type JobRepositoryImpl struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) JobRepository {
	return &JobRepositoryImpl{
		db: db,
	}
}
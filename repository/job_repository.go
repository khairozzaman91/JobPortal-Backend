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
	db      *sqlx.DB
	jobList []domain.Job
}

func NewJobRepository(db *sqlx.DB) JobRepository {
	repo := &JobRepositoryImpl{
		db: db,
	}

	GenerateInitPost(repo)

	return repo
}
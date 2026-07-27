package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

type JobSeekerRepository interface {
	Store(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error)
	List() ([]domain.JobSeekerProfile, error)
	Get(userID uint) (*domain.JobSeekerProfile, error)
	Update(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error)
	Delete(userID uint) error
}

type JobSeekerRepositoryImpl struct {
	db *sqlx.DB
}

func NewJobSeekerRepository(db *sqlx.DB) JobSeekerRepository {
	return &JobSeekerRepositoryImpl{
		db: db,
	}
}

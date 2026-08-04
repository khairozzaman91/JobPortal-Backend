package service

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type JobService interface {
	Store(job domain.Job) (domain.Job, error)
	List(page, limit int64) ([]*domain.Job, error)
	Count() (int64, error)
	Get(id int) (*domain.Job, error)
	Update(job domain.Job) (domain.Job, error)
	Delete(id uint) error
	DeleteAll() error
}
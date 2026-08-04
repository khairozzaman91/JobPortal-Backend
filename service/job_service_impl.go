package service

import (
	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/repository"
)

type JobServiceImpl struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) *JobServiceImpl {
	return &JobServiceImpl{
		repo: repo,
	}
}

func (s *JobServiceImpl) Store(job domain.Job) (domain.Job, error) {
	return s.repo.Store(job)
}

func (s *JobServiceImpl) List(page, limit int64) ([]*domain.Job, error) {
	return s.repo.List(page, limit)
}

func (s *JobServiceImpl) Count() (int64, error) {
	return s.repo.Count()
}

func (s *JobServiceImpl) Get(id int) (*domain.Job, error) {
	return s.repo.Get(id)
}

func (s *JobServiceImpl) Update(job domain.Job) (domain.Job, error) {
	return s.repo.Update(job)
}

func (s *JobServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *JobServiceImpl) DeleteAll() error {
	return s.repo.DeleteAll()
}
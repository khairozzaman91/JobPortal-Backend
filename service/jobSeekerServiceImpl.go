package service

import (
	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/repository"
)

type JobSeekerServiceImpl struct {
	repo repository.JobSeekerRepository
}

func NewJobSeekerService(repo repository.JobSeekerRepository) JobSeekerService {
	return &JobSeekerServiceImpl{
		repo: repo,
	}
}

func (s *JobSeekerServiceImpl) Store(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error) {
	return s.repo.Store(profile)
}

func (s *JobSeekerServiceImpl) List() ([]domain.JobSeekerProfile, error) {
	return s.repo.List()
}
func (s *JobSeekerServiceImpl) Get(userID uint) (*domain.JobSeekerProfile, error) {
	return s.repo.Get(userID)
}

func (s *JobSeekerServiceImpl) Update(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error) {
	return s.repo.Update(profile)
}

func (s *JobSeekerServiceImpl) Delete(userID uint) error {
	return s.repo.Delete(userID)
}

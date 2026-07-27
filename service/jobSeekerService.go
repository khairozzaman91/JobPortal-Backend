package service

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type JobSeekerService interface {
	Store(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error)
	List() ([]domain.JobSeekerProfile, error)
	Get(userID uint) (*domain.JobSeekerProfile, error)
	Update(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error)
	Delete(userID uint) error
}

package infra

import "github.com/khairozzaman91/JobPortal-Backend/domain"

type JobRepository struct {
	jobList []domain.Job
}

func NewJobRepository() *JobRepository {
	repo := &JobRepository{}

	GenerateInitPost(repo)

	return repo
}
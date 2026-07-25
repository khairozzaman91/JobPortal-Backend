package jobs

import(
	"github.com/khairozzaman91/JobPortal-Backend/repository"
	 middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

type JobHandler struct {
	repo        repository.JobRepository
	middlewares *middlewares.AuthMiddleware
}

func NewJobHandler(repo repository.JobRepository, middlewares *middlewares.AuthMiddleware) *JobHandler {
	return &JobHandler{
        repo: repo,
		middlewares: middlewares,
	}
}

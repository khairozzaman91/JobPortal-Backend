package jobseeker

import (
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/service"
)

type JobSeekerHandler struct {
	service     service.JobSeekerService
	middlewares *middlewares.AuthMiddleware
}

func NewJobSeekerHandler(
	service service.JobSeekerService,
	middlewares *middlewares.AuthMiddleware,
) *JobSeekerHandler {

	return &JobSeekerHandler{
		service:     service,
		middlewares: middlewares,
	}
}
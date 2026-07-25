package jobs

import (
	"github.com/khairozzaman91/JobPortal-Backend/service"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

type JobHandler struct {
	service     service.JobService
	middlewares *middlewares.AuthMiddleware
}

func NewJobHandler(service service.JobService, middlewares *middlewares.AuthMiddleware) *JobHandler {
	return &JobHandler{
		service:     service,
		middlewares: middlewares,
	}
}
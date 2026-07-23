package jobs

import middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"

type JobHandler struct {
	middlewares *middlewares.AuthMiddleware
}

func NewJobHandler(middlewares *middlewares.AuthMiddleware) *JobHandler {
	return &JobHandler{
		middlewares: middlewares,
	}
}

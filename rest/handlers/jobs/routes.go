package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func(h *JobHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"GET /jobs",
		http.HandlerFunc(h.GetJobs),
	)

	mux.Handle(
		"POST /jobs",
		manager.With(
			http.HandlerFunc(h.CreatePost),
			middlewares.Authorization,
			middlewares.RequireRole("admin", "employer"),
			
		),
	)

	mux.Handle(
		"PUT /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.UpdatePost),
			middlewares.Authorization,
			middlewares.RequireRole("admin", "employer"),
		),
	)

	mux.Handle(
		"DELETE /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.DeletePost),
			middlewares.Authorization,
			middlewares.RequireRole("admin", "employer"),
		),
	)
}
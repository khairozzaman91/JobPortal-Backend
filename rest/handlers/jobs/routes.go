package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *JobHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"GET /jobs",
		http.HandlerFunc(h.GetJobs),
	)

	mux.Handle(
		"POST /jobs",
		manager.With(
			http.HandlerFunc(h.CreatePost),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin", "employer"),
		),
	)

	mux.Handle(
		"PUT /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.UpdatePost),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin", "employer"),
		),
	)

	mux.Handle(
		"DELETE /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.DeletePost),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin", "employer"),
		),
	)
}

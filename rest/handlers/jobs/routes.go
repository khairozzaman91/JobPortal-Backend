package jobs

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *JobHandler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middlewares.Manager,
	rateLimiter *middlewares.RateLimiter,
) {

	mux.Handle(
		"GET /jobs",
		manager.With(
			http.HandlerFunc(h.GetJobs),
			rateLimiter.Limit,
		),
	)

	mux.Handle(
		"GET /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.GetById),
			rateLimiter.Limit,
		),
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
		"PATCH /jobs/{id}",
		manager.With(
			http.HandlerFunc(h.PatchPost),
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

	mux.Handle(
		"DELETE /jobs",
		manager.With(
			http.HandlerFunc(h.DeleteAllPosts),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)
}
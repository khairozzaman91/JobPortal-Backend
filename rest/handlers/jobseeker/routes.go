package jobseeker

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *JobSeekerHandler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middlewares.Manager,
	rateLimiter *middlewares.RateLimiter,
) {

	mux.Handle(
		"POST /jobseeker/profile",
		manager.With(
			http.HandlerFunc(h.CreateProfile),
			rateLimiter.Limit,
			h.middlewares.Authorization,
		),
	)

	mux.Handle(
		"GET /jobseeker/profile",
		manager.With(
			http.HandlerFunc(h.GetProfile),
			h.middlewares.Authorization,
		),
	)

	mux.Handle(
		"PUT /jobseeker/profile",
		manager.With(
			http.HandlerFunc(h.UpdateProfile),
			rateLimiter.Limit,
			h.middlewares.Authorization,
		),
	)

	mux.Handle(
		"DELETE /jobseeker/profile",
		manager.With(
			http.HandlerFunc(h.DeleteProfile),
			h.middlewares.Authorization,
		),
	)
}
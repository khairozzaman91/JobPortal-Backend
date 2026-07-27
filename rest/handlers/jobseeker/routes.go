package jobseeker

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *JobSeekerHandler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middlewares.Manager,
) {

	mux.Handle(
		"POST /jobseeker/profile",
		manager.With(
			http.HandlerFunc(h.CreateProfile),
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
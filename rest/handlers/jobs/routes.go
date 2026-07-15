package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"GET /jobs",
		http.HandlerFunc(GetJobs),
	)

	mux.Handle(
		"POST /jobs",
		manager.With(
			http.HandlerFunc(CreatePost),
			middlewares.Authorization,
		),
	)

	mux.Handle(
		"PUT /jobs/{id}",
		manager.With(
			http.HandlerFunc(UpdatePost),
			middlewares.Authorization,
		),
	)

	mux.Handle(
		"DELETE /jobs/{id}",
		manager.With(
			http.HandlerFunc(DeletePost),
			middlewares.Authorization,
		),
	)
}
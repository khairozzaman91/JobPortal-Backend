package user

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /users",
		http.HandlerFunc(h.CreateUser),
	)

	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(h.GetUsers),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"POST /login",
		http.HandlerFunc(h.LoginUser),
	)
}

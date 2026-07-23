package user

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
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
			middlewares.Authorization,
			middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"POST /login",
		http.HandlerFunc(h.LoginUser),
	)
}

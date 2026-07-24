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


	mux.Handle(
		"GET /users/{id}",
		manager.With(
			http.HandlerFunc(h.GetUserByID),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"PUT /users/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateUser),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"PATCH /users/{id}",
		manager.With(
			http.HandlerFunc(h.PatchUser),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"DELETE /users/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteUser),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)

	mux.Handle(
		"DELETE /users",
		manager.With(
			http.HandlerFunc(h.DeleteAllUsers),
			h.middlewares.Authorization,
			h.middlewares.RequireRole("admin"),
		),
	)
}

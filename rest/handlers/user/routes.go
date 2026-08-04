package user

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func (h *UserHandler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middlewares.Manager,
	rateLimiter *middlewares.RateLimiter,
) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
			rateLimiter.Limit,
		),
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
		manager.With(
			http.HandlerFunc(h.LoginUser),
			rateLimiter.Limit,
		),
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
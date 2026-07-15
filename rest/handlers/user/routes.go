package user

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /users",
		http.HandlerFunc(CreateUser),
	)

	mux.Handle(
		"GET /users",
		http.HandlerFunc(GetUsers),
	)

	mux.Handle(
		"POST /login",
		http.HandlerFunc(LoginUser),
	)
}
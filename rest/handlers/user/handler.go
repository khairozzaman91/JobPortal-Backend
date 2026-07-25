package user

import (
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/service"
)

type UserHandler struct {
	service     service.UserService
	middlewares *middlewares.AuthMiddleware
}

func NewUserHandler(service service.UserService, middlewares *middlewares.AuthMiddleware) *UserHandler {
	return &UserHandler{
		service:     service,
		middlewares: middlewares,
	}
}

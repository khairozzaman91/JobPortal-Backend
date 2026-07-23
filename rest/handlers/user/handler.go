package user

import "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"

type UserHandler struct {
	middlewares middlewares.AuthMiddleware
}

func NewUserHandler(middlewares *middlewares.AuthMiddleware) *UserHandler {
	return &UserHandler{
        middlewares: *middlewares,
	}
}
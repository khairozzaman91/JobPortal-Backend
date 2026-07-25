package user

import (
	"github.com/khairozzaman91/JobPortal-Backend/repository"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

type UserHandler struct {
	repo        repository.UserRepository
	middlewares *middlewares.AuthMiddleware
}

func NewUserHandler(repo repository.UserRepository, middlewares *middlewares.AuthMiddleware) *UserHandler {
	return &UserHandler{
		repo:        repo,
		middlewares: middlewares,
	}
}
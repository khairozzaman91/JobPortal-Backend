package middlewares

import "github.com/khairozzaman91/JobPortal-Backend/config"

type AuthMiddleware struct {
	cnf *config.Config
}

func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		cnf: cfg,
	}
}
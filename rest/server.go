package rest

import (
	"github.com/khairozzaman91/JobPortal-Backend/cmd"
	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

func Server() {
	cnf := config.GetConfig()

	authMiddleware := middlewares.NewAuthMiddleware(cnf)
	jobHandler := jobs.NewJobHandler(authMiddleware)
	userHandler := user.NewUserHandler(authMiddleware)

	server := cmd.NewServer(cnf, jobHandler, userHandler)
	server.Start()
}

package rest

import (
	"github.com/khairozzaman91/JobPortal-Backend/cmd"
	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/infra"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/service"
)

func Server() {
	cnf := config.GetConfig()

	authMiddleware := middlewares.NewAuthMiddleware(cnf)

	jobRepo := infra.NewJobRepository()
	userRepo := infra.NewUserRepository()

	jobService := service.NewJobService(jobRepo)
	
	jobHandler := jobs.NewJobHandler(jobService, authMiddleware)
	userHandler := user.NewUserHandler(userRepo, authMiddleware)

	server := cmd.NewServer(cnf, jobHandler, userHandler)
	server.Start()
}

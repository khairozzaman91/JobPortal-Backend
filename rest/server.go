package rest

import (
	"log"

	"github.com/khairozzaman91/JobPortal-Backend/cmd"
	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/infra/postgres"
	"github.com/khairozzaman91/JobPortal-Backend/repository"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobseeker"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/service"
)

func Server() {
	cnf := config.GetConfig()

	// Database Connection
	db, err := postgres.NewConnection(cnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authMiddleware := middlewares.NewAuthMiddleware(cnf)

	// Repository
	jobRepo := repository.NewJobRepository(db)
	userRepo := repository.NewUserRepository(db)
	jobSeekerRepo := repository.NewJobSeekerRepository(db)

	// Service
	jobService := service.NewJobService(jobRepo)
	userService := service.NewUserService(userRepo)
	jobSeekerService := service.NewJobSeekerService(jobSeekerRepo)

	// Handler
	jobHandler := jobs.NewJobHandler(jobService, authMiddleware)
	userHandler := user.NewUserHandler(userService, authMiddleware)
	jobSeekerHandler := jobseeker.NewJobSeekerHandler(jobSeekerService, authMiddleware)

	// Server
	server := cmd.NewServer(
		cnf,
		jobHandler,
		userHandler,
		jobSeekerHandler,
	)

	server.Start()
}

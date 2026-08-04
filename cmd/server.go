package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobseeker"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

type Server struct {
	cnf              *config.Config
	jobHandler       *jobs.JobHandler
	userHandler      *user.UserHandler
	jobSeekerHandler *jobseeker.JobSeekerHandler
}

func NewServer(
	cnf *config.Config,
	jobHandler *jobs.JobHandler,
	userHandler *user.UserHandler,
	jobSeekerHandler *jobseeker.JobSeekerHandler,
) *Server {
	return &Server{
		cnf:              cnf,
		jobHandler:       jobHandler,
		userHandler:      userHandler,
		jobSeekerHandler: jobSeekerHandler,
	}
}

func (server *Server) Start() {

	manager := middlewares.NewManager()
	rateLimiter := middlewares.NewRateLimiter(
		5,
		time.Minute,
	)

	mux := http.NewServeMux()

	manager.Use(
		middlewares.CORSMiddleware,
		middlewares.Logger,
	)

	// Register Routes
	server.jobHandler.RegisterRoutes(
		mux,
		manager,
		rateLimiter,
	)

	server.userHandler.RegisterRoutes(
		mux,
		manager,
		rateLimiter,
	)

	server.jobSeekerHandler.RegisterRoutes(
		mux,
		manager,
		rateLimiter,
	)

	handler := manager.Wrapper(mux)

	port := fmt.Sprintf(":%d", server.cnf.HTTPPort)

	fmt.Println("Server Running on", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

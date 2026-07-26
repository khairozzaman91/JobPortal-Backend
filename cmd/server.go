package cmd

import (
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
)

type Server struct {
	cnf         *config.Config
	jobHandler  *jobs.JobHandler
	userHandler *user.UserHandler
}

func NewServer(
	cnf *config.Config,
	jobHandler *jobs.JobHandler,
	userHandler *user.UserHandler,
) *Server {
	return &Server{
		cnf:         cnf,
		jobHandler:  jobHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {

	manager := middlewares.NewManager()

	mux := http.NewServeMux()

	manager.Use(
		middlewares.CORSMiddleware,
		middlewares.Logger,
	)

	// Register Routes
	server.jobHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	handler := manager.Wrapper(mux)

	port := fmt.Sprintf(":%d", server.cnf.HTTPPort)

	fmt.Println("Server Running on", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/infra/postgres"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func Server() {

	// Database Connection
	db, err := postgres.GetConnect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	
	manager := middlewares.NewManager()

	mux := http.NewServeMux()

	cfg := config.GetConfig()

	manager.Use(
		middlewares.CORSMiddleware,
		middlewares.Logger,
	)

	// Register Routes
	jobs.RegisterRoutes(mux, manager)
	user.RegisterRoutes(mux, manager)

	handler := manager.Wraper(mux)

	port := fmt.Sprintf(":%d", cfg.HTTPPort)

	fmt.Println("Server Running on", port)

	err = http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}

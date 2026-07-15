package cmd

import (
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func Server() {

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

	fmt.Println("Server Running on port :8080")

	port := fmt.Sprintf(":%d", cfg.HTTPPort)

	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}

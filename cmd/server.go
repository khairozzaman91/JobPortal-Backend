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

	mux := http.NewServeMux()

	cg := config.GetConfig()

	// job post  handler
	mux.Handle("GET /jobs",middlewares.Logger(http.HandlerFunc(jobs.GetJobs)))
	mux.Handle("POST /jobs", middlewares.Logger(middlewares.Authorization(http.HandlerFunc(jobs.CreatePost))))
	mux.Handle("PUT /jobs/{id}",middlewares.Logger(middlewares.Authorization(http.HandlerFunc(jobs.UpdatePost))))
	mux.Handle("DELETE /jobs/{id}",middlewares.Logger(middlewares.Authorization(http.HandlerFunc(jobs.DeletePost))))

	// user handler
	mux.Handle("POST /users",middlewares.Logger(http.HandlerFunc(user.CreateUser)))
	mux.Handle("GET /users", middlewares.Logger(http.HandlerFunc(user.GetUsers)))
	mux.Handle("POST /login", middlewares.Logger(http.HandlerFunc(user.LoginUser)))



	fmt.Println("Server Running on port : 8080")

	handler := middlewares.CORSMiddleware(mux)
	port := fmt.Sprintf(":%d", cg.HTTPPort)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}


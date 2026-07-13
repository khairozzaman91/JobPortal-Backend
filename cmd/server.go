package cmd

import (
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func Server() {

	mux := http.NewServeMux()

	// job post  handler
	mux.Handle("GET /jobs", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(jobs.GetJobs))))
	mux.Handle("POST /jobs", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(jobs.CreatePost))))
	mux.Handle("PUT /jobs/{id}", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(jobs.UpdatePost))))
	mux.Handle("DELETE /jobs/{id}", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(jobs.DeletePost))))

	// user handler
	mux.Handle("POST /users", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(user.CreateUser))))
	mux.Handle("GET /users", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(user.GetUsers))))
	mux.Handle("POST /login", middlewares.CORSMiddleware(middlewares.Logger(http.HandlerFunc(user.LoginUser))))



	fmt.Println("Server Running on port : 3000")

	handler := middlewares.CORSMiddleware(mux)
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}

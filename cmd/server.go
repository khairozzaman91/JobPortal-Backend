package cmd

import (
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
)

func Server() {

	mux := http.NewServeMux()

	mux.Handle("GET /jobs", middlewares.Logger(http.HandlerFunc(jobs.GetJobs)))
	mux.Handle("POST /jobs", middlewares.Logger(http.HandlerFunc(jobs.CreatePost)))
	mux.Handle("PUT /jobs/{id}", middlewares.Logger(http.HandlerFunc(jobs.UpdatePost)))
	mux.Handle("DELETE /jobs/{id}",middlewares.Logger(http.HandlerFunc(jobs.DeletePost)))

	fmt.Println("Server Running on port : 3000")

	handler := middlewares.CORSMiddleware(mux)
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}

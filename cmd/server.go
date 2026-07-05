package cmd

import (
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
)

func Server() {

	mux := http.NewServeMux()

	mux.Handle("GET /jobs", http.HandlerFunc(jobs.GetJobs))

	fmt.Println("Server Running on port : 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		return
	}
}

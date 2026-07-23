package main

import (
	"github.com/khairozzaman91/JobPortal-Backend/cmd"
	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/jobs"
	"github.com/khairozzaman91/JobPortal-Backend/rest/handlers/user"
)

func main() {
	cnf := config.GetConfig()

	jobHandler := jobs.NewJobHandler()
	userHandler := user.NewUserHandler()
	server := cmd.NewServer(*cnf, jobHandler, userHandler)
	server.Start()
}
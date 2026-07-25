package jobs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) PatchPost(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid job id")
		return
	}

	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	job, err := h.repo.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	if job.PostedBy != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only update your own jobs")
		return
	}

	var req domain.Job

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	if req.Title != "" {
		job.Title = req.Title
	}

	if req.Description != "" {
		job.Description = req.Description
	}

	if req.CompanyName != "" {
		job.CompanyName = req.CompanyName
	}

	if req.Location != "" {
		job.Location = req.Location
	}

	if req.Salary != 0 {
		job.Salary = req.Salary
	}

	if req.JobType != "" {
		job.JobType = req.JobType
	}

	if req.ExperienceLevel != "" {
		job.ExperienceLevel = req.ExperienceLevel
	}

	updatedJob, err := h.repo.Update(*job)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	utils.SendData(w, updatedJob, http.StatusOK)
}
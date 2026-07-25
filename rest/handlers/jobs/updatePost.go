package jobs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid job id")
		return
	}

	// Logged-in user
	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Existing job
	job, err := h.repo.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	// Ownership check
	if job.PostedBy != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only update your own jobs")
		return
	}

	var updatedJob domain.Job

	if err := json.NewDecoder(r.Body).Decode(&updatedJob); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Preserve immutable fields
	updatedJob.ID = job.ID
	updatedJob.PostedBy = job.PostedBy

	updated, err := h.repo.Update(updatedJob)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	utils.SendData(w, updated, http.StatusOK)
}

package jobs

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	strId := r.PathValue("id")
	if strId == "" {
		utils.SendError(w, http.StatusBadRequest, "Missing ID")
		return
	}

	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid job ID")
		return
	}

	// Get logged-in user
	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var newList []domain.Job
	found := false

	for _, job := range dto.JobList {

		if job.ID == uint(id) {
			found = true

			// Ownership Check
			if claims.Role != "admin" && job.PostedBy != uint(claims.Sub) {
				utils.SendError(w, http.StatusForbidden, "You can only delete your own jobs")
				return
			}

			// Skip this job (delete)
			continue
		}

		newList = append(newList, job)
	}

	if !found {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	// Update the original slice
	dto.JobList = newList
	utils.SendData(w, "Successfully deleted job", http.StatusOK)
}

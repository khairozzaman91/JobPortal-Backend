package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get logged-in user from Context
	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var newJob domain.Job

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newJob)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Generate Job ID
	newJob.ID = uint(len(dto.JobList) + 1)

	// Set Owner
	newJob.PostedBy = uint(claims.Sub)

	// Save Job
	dto.JobList = append(dto.JobList, newJob)

	// Success Response
	utils.SendData(w, newJob, http.StatusCreated)
}

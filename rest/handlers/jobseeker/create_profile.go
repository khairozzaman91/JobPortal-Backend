package jobseeker

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobSeekerHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {

	var profile domain.JobSeekerProfile

	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	profile, err = h.service.Store(profile)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendData(w, profile, http.StatusCreated)
}
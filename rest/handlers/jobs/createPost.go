package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var newJob domain.Job

	if err := json.NewDecoder(r.Body).Decode(&newJob); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Set Owner
	newJob.PostedBy = uint(claims.Sub)

	// Save Job
	job, err := h.service.Store(newJob)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendData(w, job, http.StatusCreated)
}
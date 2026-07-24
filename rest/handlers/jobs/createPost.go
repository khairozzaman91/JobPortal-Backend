package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/infra"
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
	infra.Store(newJob)

	utils.SendData(w, newJob, http.StatusCreated)
}

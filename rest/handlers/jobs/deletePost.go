package jobs

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/infra"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid job ID")
		return
	}

	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	job := infra.Get(id)
	if job == nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	// Ownership Check
	if claims.Role != "admin" && job.PostedBy != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only delete your own jobs")
		return
	}

	infra.Delete(job.ID)

	utils.SendData(w, "Successfully deleted job", http.StatusOK)
}

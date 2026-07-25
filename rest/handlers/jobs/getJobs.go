package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.repo.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendData(w, jobs, http.StatusOK)
}
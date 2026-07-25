package jobs

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid job id")
		return
	}

	job, err := h.service.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Job not found")
		return
	}

	utils.SendData(w, job, http.StatusOK)
}
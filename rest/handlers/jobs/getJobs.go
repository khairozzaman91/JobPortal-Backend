package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	utils.SendData(w, dto.JobList, http.StatusOK)
}

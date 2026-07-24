package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/infra"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, infra.List(), http.StatusOK)
}

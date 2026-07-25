package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) DeleteAllPosts(w http.ResponseWriter, r *http.Request) {

	err := h.repo.DeleteAll()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendData(w, "All jobs deleted successfully", http.StatusOK)
}
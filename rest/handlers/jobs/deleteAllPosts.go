package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/infra"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobHandler) DeleteAllPosts(w http.ResponseWriter, r *http.Request) {

	infra.DeleteAll()

	utils.SendData(w, "All jobs deleted successfully", http.StatusOK)
}

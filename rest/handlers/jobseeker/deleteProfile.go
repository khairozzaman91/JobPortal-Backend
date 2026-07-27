package jobseeker

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobSeekerHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("user").(middlewares.Claims)

	err := h.service.Delete(uint(claims.Sub))
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendData(w, "Profile deleted successfully", http.StatusOK)
}
package jobseeker

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *JobSeekerHandler) GetProfile(w http.ResponseWriter, r *http.Request) {

	claims := r.Context().Value("user").(middlewares.Claims)

	profile, err := h.service.Get(uint(claims.Sub))
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Profile not found")
		return
	}

	utils.SendData(w, profile, http.StatusOK)
}
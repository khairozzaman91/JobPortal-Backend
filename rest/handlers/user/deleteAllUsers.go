package user

import (
	"net/http"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if claims.Role != "admin" {
		utils.SendError(w, http.StatusForbidden, "Only admin can delete all users")
		return
	}

	err := h.service.DeleteAll()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete all users")
		return
	}

	utils.SendData(w, "All users deleted successfully", http.StatusOK)
}

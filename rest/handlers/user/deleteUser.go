package user

import (
	"net/http"
	"strconv"

	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.repo.Get(id)
	if err != nil || user == nil {
		utils.SendError(w, http.StatusNotFound, "User not found")
		return
	}

	if claims.Role != "admin" && user.ID != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only delete your own profile")
		return
	}

	err = h.repo.Delete(user.ID)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.SendData(w, "Successfully deleted user", http.StatusOK)
}
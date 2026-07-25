package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := h.service.Get(id)
	if err != nil || user == nil {
		utils.SendError(w, http.StatusNotFound, "User not found")
		return
	}

	// Authorization Check
	if claims.Role != "admin" && user.ID != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only update your own profile")
		return
	}

	var updatedUser domain.User

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Preserve immutable fields
	updatedUser.ID = user.ID
	updatedUser.CreatedAt = user.CreatedAt
	updatedUser.UpdatedAt = time.Now()

	updatedUser, err = h.service.Update(updatedUser)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	utils.SendData(w, updatedUser, http.StatusOK)
}
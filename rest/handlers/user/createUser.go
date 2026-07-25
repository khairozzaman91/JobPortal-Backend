package user

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Role Validation
	switch user.Role {
	case "admin", "employer", "jobseeker":
	default:
		utils.SendError(w, http.StatusBadRequest, "Invalid role")
		return
	}

	user, err := h.service.Store(user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendData(w, user, http.StatusCreated)
}

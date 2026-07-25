package user

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	utils.SendData(w, users, http.StatusOK)
}
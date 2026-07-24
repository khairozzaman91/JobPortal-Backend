package user

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/infra"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	user := infra.UserGet(id)
	if user == nil {
		utils.SendError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.SendData(w, user, http.StatusOK)
}

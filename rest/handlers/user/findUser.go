package user

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/infra"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, infra.UserList(), http.StatusOK)
}

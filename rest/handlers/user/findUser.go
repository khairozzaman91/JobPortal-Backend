package user

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	utils.SendData(w, domain.UserList, http.StatusOK)
}

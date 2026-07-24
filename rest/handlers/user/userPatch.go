package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/infra"
	middlewares "github.com/khairozzaman91/JobPortal-Backend/rest/middleware"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func (h *UserHandler) PatchUser(w http.ResponseWriter, r *http.Request) {

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

	user := infra.UserGet(id)
	if user == nil {
		utils.SendError(w, http.StatusNotFound, "User not found")
		return
	}

	if claims.Role != "admin" && user.ID != uint(claims.Sub) {
		utils.SendError(w, http.StatusForbidden, "You can only update your own profile")
		return
	}

	var req domain.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	user.UpdatedAt = time.Now()

	infra.UserUpdate(*user)

	utils.SendData(w, user, http.StatusOK)
}
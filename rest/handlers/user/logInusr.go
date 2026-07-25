package user

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	cfg := config.GetConfig()

	users, err := h.repo.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	for _, user := range users {

		if user.Email == req.Email && user.Password == req.Password {

			claims := utils.Payload{
				Sub:       int(user.ID),
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Role:      user.Role,
			}

			token, err := utils.CreateJwt(cfg.JWTSecret, claims)
			if err != nil {
				utils.SendError(w, http.StatusInternalServerError, "Failed to generate token")
				return
			}

			utils.SendData(w, token, http.StatusOK)
			return
		}
	}

	utils.SendError(w, http.StatusUnauthorized, "Invalid email or password")
}

package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Handler Called")
	var req LoginRequest

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&req)

	if err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}

	cfg := config.GetConfig()

	for _, user := range domain.UserList {

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
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}

			utils.SendData(w, token, http.StatusOK)
			return
		}
	}

	utils.SendError(w, http.StatusUnauthorized, "Invalid email or password")
}

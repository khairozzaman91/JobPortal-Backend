package user

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&req)

	if err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}

	for _, user := range domain.UserList {

		if user.Email == req.Email && user.Password == req.Password {
			utils.SendData(w, "Login successful", http.StatusOK)
			return
		}
	}

	utils.SendError(w, http.StatusUnauthorized, "Invalid email or password")

}

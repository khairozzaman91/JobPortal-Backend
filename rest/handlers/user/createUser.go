package user

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}

	// ID Generate
	user.ID = uint(len(domain.UserList) + 1)

	// Append to list
	domain.UserList = append(domain.UserList, &user)

	utils.SendData(w, user, http.StatusCreated)
}

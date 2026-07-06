package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	// Cross check
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newJob domain.Job

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newJob)
	if err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}

	// Append to jobList

	newJob.ID = uint(len(dto.JobList) + 1)
	dto.JobList = append(dto.JobList, newJob)

	// Response
	encode := json.NewEncoder(w)
	encode.Encode(newJob)

	http.Error(w, "Create sucessfully", http.StatusOK)
}

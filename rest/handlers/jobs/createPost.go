package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

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
	w.WriteHeader(http.StatusOK)
	encode.Encode(newJob)

	http.Error(w, "Create sucessfully", http.StatusOK)
}

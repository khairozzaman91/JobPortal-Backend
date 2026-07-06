package jobs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	strId := r.PathValue("id")
	if strId == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	var updatelist domain.Job

	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&updatelist)
	if err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}
	for i := range dto.JobList {
		if dto.JobList[i].ID == uint(id) {
			dto.JobList[i].Title = updatelist.Title
			dto.JobList[i].Description = updatelist.Description
			dto.JobList[i].CompanyName = updatelist.CompanyName
			dto.JobList[i].Location = updatelist.Location
			dto.JobList[i].Salary = updatelist.Salary
			dto.JobList[i].JobType = updatelist.JobType
			dto.JobList[i].ExperienceLevel = updatelist.ExperienceLevel

			encode := json.NewEncoder(w)
			w.WriteHeader(http.StatusOK)
			encode.Encode(dto.JobList[i])
			return
		}
	}
	http.Error(w, "Job not found", http.StatusNotFound)
}

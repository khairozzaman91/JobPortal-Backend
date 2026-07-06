package jobs

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
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

	var newList []domain.Job

	for _, job := range dto.JobList {
		if job.ID != uint(id) {
			newList = append(newList, job)
		}
	}

	dto.JobList = newList

	utils.SendData(w, "successfully Deleted Job", http.StatusOK)
}

package jobs

import (
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	utils.SendData(w, dto.JobList, http.StatusOK)
}

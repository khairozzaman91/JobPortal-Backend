package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var newJob domain.Job

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&newJob)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Give me valid json")
		return
	}

	// Generate Job ID
	newJob.ID = uint(len(dto.JobList) + 1)

	// Save Job
	dto.JobList = append(dto.JobList, newJob)

	// Success Response
	utils.SendData(w, newJob, http.StatusCreated)
}
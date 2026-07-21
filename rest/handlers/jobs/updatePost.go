package jobs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/dto"
	"github.com/khairozzaman91/JobPortal-Backend/rest/middlewares"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
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

	// Get logged-in user
	claims, ok := r.Context().Value("user").(middlewares.Claims)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
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
			// Ownership Check
			if claims.Role != "admin" && dto.JobList[i].PostedBy != uint(claims.Sub) {
				utils.SendError(w, http.StatusForbidden, "You can only update your own jobs")
				return
			}
			dto.JobList[i].Title = updatelist.Title
			dto.JobList[i].Description = updatelist.Description
			dto.JobList[i].CompanyName = updatelist.CompanyName
			dto.JobList[i].Location = updatelist.Location
			dto.JobList[i].Salary = updatelist.Salary
			dto.JobList[i].JobType = updatelist.JobType
			dto.JobList[i].ExperienceLevel = updatelist.ExperienceLevel
			utils.SendData(w, dto.JobList[i], http.StatusOK)
			return
		}
	}

	utils.SendError(w, http.StatusNotFound, "Job not found")

}

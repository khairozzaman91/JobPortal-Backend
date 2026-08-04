package jobs

import (
	"net/http"
	"strconv"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

type Pagination struct {
	Data       []*domain.Job `json:"data"`
	Page       int64         `json:"page"`
	Limit      int64         `json:"limit"`
	TotalItems int64         `json:"total_items"`
	TotalPages int64         `json:"total_pages"`
}

func (h *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	jobslist, err := h.service.List(page, limit)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cnt, err := h.service.Count()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	paginatedData := Pagination{
		Data:       jobslist,
		Page:       page,
		Limit:      limit,
		TotalItems: cnt,
		TotalPages: cnt / limit,
	}

	utils.SendData(w, paginatedData, http.StatusOK)
}

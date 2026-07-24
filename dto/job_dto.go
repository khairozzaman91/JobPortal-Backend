package dto

type CreateJobRequest struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	CompanyName     string  `json:"company_name" validate:"required"`
	Location        string  `json:"location"`
	Salary          float64 `json:"salary"`
	JobType         string  `json:"job_type"`
	ExperienceLevel string  `json:"experience_level"`
}

type JobResponse struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	CompanyName     string  `json:"company_name"`
	Location        string  `json:"location"`
	Salary          float64 `json:"salary"`
	JobType         string  `json:"job_type"`
	ExperienceLevel string  `json:"experience_level"`
	PostedBy        uint    `json:"posted_by"`
	IsActive        bool    `json:"is_active"`
}

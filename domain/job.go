package domain

import "time"

type Job struct {
	ID              uint      `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Description     string    `json:"description" db:"description"`
	CompanyName     string    `json:"company_name" db:"company_name"`
	Location        string    `json:"location" db:"location"`
	Salary          float64   `json:"salary" db:"salary"`
	JobType         string    `json:"job_type" db:"job_type"`
	ExperienceLevel string    `json:"experience_level" db:"experience_level"`
	PostedBy        uint      `json:"posted_by" db:"posted_by"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

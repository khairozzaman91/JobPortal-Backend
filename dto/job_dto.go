package dto

import "github.com/khairozzaman91/JobPortal-Backend/domain"

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

var JobList []domain.Job

func init() {

	job1 := domain.Job{
		ID:              1,
		Title:           "Senior Software Engineer",
		Description:     "We are looking for an experienced Go developer to join our team.",
		CompanyName:     "TechCorp Ltd",
		Location:        "Dhaka",
		Salary:          120000,
		JobType:         "Full-time",
		ExperienceLevel: "Senior",
		PostedBy:        1,
		IsActive:        true,
	}
	job2 := domain.Job{
		ID:              2,
		Title:           "Junior Backend Developer",
		Description:     "Looking for a passionate backend developer with Node.js experience.",
		CompanyName:     "Innovate Solutions",
		Location:        "Chittagong",
		Salary:          45000,
		JobType:         "Full-time",
		ExperienceLevel: "Junior",
		PostedBy:        1,
		IsActive:        true,
	}

	job3 := domain.Job{
		ID:              3,
		Title:           "Frontend Developer",
		Description:     "React.js expert needed for our growing startup.",
		CompanyName:     "PixelCraft",
		Location:        "Dhaka",
		Salary:          60000,
		JobType:         "Full-time",
		ExperienceLevel: "Mid",
		PostedBy:        2,
		IsActive:        true,
	}

	job4 := domain.Job{
		ID:              4,
		Title:           "Data Analyst",
		Description:     "Data-driven decisions are key. Join our analytics team.",
		CompanyName:     "DataVision",
		Location:        "Sylhet",
		Salary:          55000,
		JobType:         "Full-time",
		ExperienceLevel: "Mid",
		PostedBy:        2,
		IsActive:        true,
	}

	job5 := domain.Job{
		ID:              5,
		Title:           "DevOps Engineer",
		Description:     "Looking for an experienced DevOps engineer with AWS knowledge.",
		CompanyName:     "CloudScale",
		Location:        "Dhaka",
		Salary:          95000,
		JobType:         "Full-time",
		ExperienceLevel: "Senior",
		PostedBy:        1,
		IsActive:        true,
	}

	job6 := domain.Job{
		ID:              6,
		Title:           "Mobile App Developer",
		Description:     "Flutter developer needed for cross-platform mobile apps.",
		CompanyName:     "AppNova",
		Location:        "Chittagong",
		Salary:          70000,
		JobType:         "Full-time",
		ExperienceLevel: "Mid",
		PostedBy:        2,
		IsActive:        true,
	}

	job7 := domain.Job{
		ID:              7,
		Title:           "UI/UX Designer",
		Description:     "Creative designer with Figma and Adobe XD experience.",
		CompanyName:     "DesignHub",
		Location:        "Dhaka",
		Salary:          65000,
		JobType:         "Full-time",
		ExperienceLevel: "Mid",
		PostedBy:        3,
		IsActive:        true,
	}

	job8 := domain.Job{
		ID:              8,
		Title:           "Marketing Manager",
		Description:     "Digital marketing expert to lead our campaigns.",
		CompanyName:     "BrandBoost",
		Location:        "Dhaka",
		Salary:          80000,
		JobType:         "Full-time",
		ExperienceLevel: "Senior",
		PostedBy:        3,
		IsActive:        true,
	}

	job9 := domain.Job{
		ID:              9,
		Title:           "Content Writer",
		Description:     "Creative content writer for our tech blog.",
		CompanyName:     "WriteFlow",
		Location:        "Sylhet",
		Salary:          40000,
		JobType:         "Full-time",
		ExperienceLevel: "Junior",
		PostedBy:        2,
		IsActive:        true,
	}

	job10 := domain.Job{
		ID:              10,
		Title:           "Project Manager",
		Description:     "Experienced PM to lead multiple software projects.",
		CompanyName:     "ProjectPro",
		Location:        "Dhaka",
		Salary:          110000,
		JobType:         "Full-time",
		ExperienceLevel: "Senior",
		PostedBy:        1,
		IsActive:        true,
	}

	JobList = append(JobList, job1)
	JobList = append(JobList, job2)
	JobList = append(JobList, job3)
	JobList = append(JobList, job4)
	JobList = append(JobList, job5)
	JobList = append(JobList, job6)
	JobList= append(JobList, job7)
	JobList= append(JobList, job8)
	JobList = append(JobList, job9)
	JobList = append(JobList, job10)
}

package repository

import (
	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

func (r *JobRepositoryImpl) Store(j domain.Job) (domain.Job, error) {
	j.ID = uint(len(r.jobList) + 1)
	r.jobList = append(r.jobList, j)
	return j, nil
}

func (r *JobRepositoryImpl) List() ([]domain.Job, error) {
	return r.jobList, nil
}

func (r *JobRepositoryImpl) Get(jobID int) (*domain.Job, error) {
	// TODO: Refactor to return original slice element using index.
	// Current implementation returns pointer to range variable.
	for _, job := range r.jobList {
		if job.ID == uint(jobID) {
			return &job, nil
		}
	}

	return nil, nil
}

func (r *JobRepositoryImpl) Update(job domain.Job) (domain.Job, error) {
	for idx, j := range r.jobList {
		if j.ID == job.ID {
			r.jobList[idx] = job
			return job, nil
		}
	}

	return domain.Job{}, nil
}

func (r *JobRepositoryImpl) Delete(jobID uint) error {
	var tempList []domain.Job

	for _, job := range r.jobList {
		if job.ID != jobID {
			tempList = append(tempList, job)
		}
	}

	r.jobList = tempList
	return nil
}

func (r *JobRepositoryImpl) DeleteAll() error {
	r.jobList = nil
	return nil
}

func GenerateInitPost(r *JobRepositoryImpl) {

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

	r.jobList = append(r.jobList,
		job1, job2, job3, job4, job5,
		job6, job7, job8, job9, job10,
	)
}
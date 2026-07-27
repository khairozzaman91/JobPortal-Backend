package domain

import "time"

type JobSeekerProfile struct {
	ID     uint `db:"id" json:"id"`
	UserID uint `db:"user_id" json:"user_id"`

	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Gender      string `db:"gender" json:"gender"`

	Address string `db:"address" json:"address"`
	Bio     string `db:"bio" json:"bio"`

	Skills     string `db:"skills" json:"skills"`
	Experience string `db:"experience" json:"experience"`
	Education  string `db:"education" json:"education"`

	LinkedinURL  string `db:"linkedin_url" json:"linkedin_url"`
	GithubURL    string `db:"github_url" json:"github_url"`
	PortfolioURL string `db:"portfolio_url" json:"portfolio_url"`

	ProfileImageURL string `db:"profile_image_url" json:"profile_image_url"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

package repository

import (
	"fmt"

	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

func (r *JobSeekerRepositoryImpl) Store(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error) {

	fmt.Println("UserID:", profile.UserID)
	query := `
		INSERT INTO job_seeker_profiles (
			user_id,
			date_of_birth,
			gender,
			address,
			bio,
			skills,
			experience,
			education,
			linkedin_url,
			github_url,
			portfolio_url,
			profile_image_url
		)
		VALUES (
			:user_id,
			:date_of_birth,
			:gender,
			:address,
			:bio,
			:skills,
			:experience,
			:education,
			:linkedin_url,
			:github_url,
			:portfolio_url,
			:profile_image_url
		)
		RETURNING id, created_at, updated_at;
	`

	rows, err := r.db.NamedQuery(query, profile)
	if err != nil {
		return domain.JobSeekerProfile{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&profile.ID,
			&profile.CreatedAt,
			&profile.UpdatedAt,
		)
		if err != nil {
			return domain.JobSeekerProfile{}, err
		}
	}

	return profile, nil
}

func (r *JobSeekerRepositoryImpl) List() ([]domain.JobSeekerProfile, error) {

	query := `
		SELECT
			id,
			user_id,
			date_of_birth,
			gender,
			address,
			bio,
			skills,
			experience,
			education,
			linkedin_url,
			github_url,
			portfolio_url,
			profile_image_url,
			created_at,
			updated_at
		FROM job_seeker_profiles
		ORDER BY id;
	`

	var profiles []domain.JobSeekerProfile

	err := r.db.Select(&profiles, query)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *JobSeekerRepositoryImpl) Get(userID uint) (*domain.JobSeekerProfile, error) {

	query := `
		SELECT
			id,
			user_id,
			date_of_birth,
			gender,
			address,
			bio,
			skills,
			experience,
			education,
			linkedin_url,
			github_url,
			portfolio_url,
			profile_image_url,
			created_at,
			updated_at
		FROM job_seeker_profiles
		WHERE user_id = $1;
	`

	var profile domain.JobSeekerProfile

	err := r.db.Get(&profile, query, userID)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *JobSeekerRepositoryImpl) Update(profile domain.JobSeekerProfile) (domain.JobSeekerProfile, error) {

	query := `
		UPDATE job_seeker_profiles
		SET
			date_of_birth = :date_of_birth,
			gender = :gender,
			address = :address,
			bio = :bio,
			skills = :skills,
			experience = :experience,
			education = :education,
			linkedin_url = :linkedin_url,
			github_url = :github_url,
			portfolio_url = :portfolio_url,
			profile_image_url = :profile_image_url,
			updated_at = NOW()
		WHERE user_id = :user_id
		RETURNING updated_at;
	`

	rows, err := r.db.NamedQuery(query, profile)
	if err != nil {
		return domain.JobSeekerProfile{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&profile.UpdatedAt)
		if err != nil {
			return domain.JobSeekerProfile{}, err
		}
	}

	return profile, nil
}

func (r *JobSeekerRepositoryImpl) Delete(userID uint) error {

	query := `
		DELETE FROM job_seeker_profiles
		WHERE user_id = $1;
	`

	_, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

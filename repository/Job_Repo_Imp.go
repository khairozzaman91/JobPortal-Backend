package repository

import (
	"github.com/khairozzaman91/JobPortal-Backend/domain"
)

func (r *JobRepositoryImpl) Store(job domain.Job) (domain.Job, error) {

	query := `
		INSERT INTO jobs (
			title,
			description,
			company_name,
			location,
			salary,
			job_type,
			experience_level,
			posted_by,
			is_active
		)
		VALUES (
			:title,
			:description,
			:company_name,
			:location,
			:salary,
			:job_type,
			:experience_level,
			:posted_by,
			:is_active
		)
		RETURNING id, created_at, updated_at;
	`

	rows, err := r.db.NamedQuery(query, job)
	if err != nil {
		return domain.Job{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&job.ID,
			&job.CreatedAt,
			&job.UpdatedAt,
		); err != nil {
			return domain.Job{}, err
		}
	}

	return job, nil
}

func (r *JobRepositoryImpl) List(page, limit int64) ([]*domain.Job, error) {

	offset := ((page - 1) * limit + 1)

	query := `
		SELECT
			id,
			title,
			description,
			company_name,
			location,
			salary,
			job_type,
			experience_level,
			posted_by,
			is_active,
			created_at,
			updated_at
		FROM jobs
		ORDER BY id
		LIMIT $1
		OFFSET $2;
	`

	var jobs []*domain.Job

	err := r.db.Select(&jobs, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *JobRepositoryImpl) Count() (int64, error) {

	query := `
		SELECT COUNT(*)
		FROM jobs;
	`

	var count int64

	err := r.db.Get(&count, query)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *JobRepositoryImpl) Get(jobID int) (*domain.Job, error) {

	query := `
		SELECT
			id,
			title,
			description,
			company_name,
			location,
			salary,
			job_type,
			experience_level,
			posted_by,
			is_active,
			created_at,
			updated_at
		FROM jobs
		WHERE id = $1;
	`

	var job domain.Job

	err := r.db.Get(&job, query, jobID)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (r *JobRepositoryImpl) Update(job domain.Job) (domain.Job, error) {

	query := `
		UPDATE jobs
		SET
			title = :title,
			description = :description,
			company_name = :company_name,
			location = :location,
			salary = :salary,
			job_type = :job_type,
			experience_level = :experience_level,
			posted_by = :posted_by,
			is_active = :is_active,
			updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at;
	`

	rows, err := r.db.NamedQuery(query, job)
	if err != nil {
		return domain.Job{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&job.UpdatedAt); err != nil {
			return domain.Job{}, err
		}
	}

	return job, nil
}

func (r *JobRepositoryImpl) Delete(jobID uint) error {

	query := `
		DELETE FROM jobs
		WHERE id = $1;
	`

	_, err := r.db.Exec(query, jobID)
	return err
}

func (r *JobRepositoryImpl) DeleteAll() error {

	query := `
		DELETE FROM jobs;
	`

	_, err := r.db.Exec(query)
	return err
}
package repository

import "github.com/khairozzaman91/JobPortal-Backend/domain"



func (r *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			phone,
			role,
			created_at,
			updated_at
		FROM users
		WHERE email = $1;
	`

	var user domain.User

	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Store(user domain.User) (domain.User, error) {

	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			phone,
			role
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:phone,
			:role
		)
		RETURNING id, created_at, updated_at;
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return domain.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (r *UserRepositoryImpl) List() ([]domain.User, error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			phone,
			role,
			created_at,
			updated_at
		FROM users
		ORDER BY id;
	`

	var users []domain.User

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) Get(userID int) (*domain.User, error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			phone,
			role,
			created_at,
			updated_at
		FROM users
		WHERE id = $1;
	`

	var user domain.User

	err := r.db.Get(&user, query, userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {

	query := `
		UPDATE users
		SET
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			password = :password,
			phone = :phone,
			role = :role,
			updated_at = NOW()
		WHERE id = :id
		RETURNING updated_at;
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return domain.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.UpdatedAt)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(userID uint) error {

	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	_, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) DeleteAll() error {

	query := `
		DELETE FROM users;
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

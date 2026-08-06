-- +migrate Up

CREATE TABLE IF NOT EXISTS job_seeker_profiles (
    id SERIAL PRIMARY KEY,

    user_id INTEGER NOT NULL UNIQUE,

    date_of_birth DATE,
    gender VARCHAR(20),

    address TEXT,
    bio TEXT,

    skills TEXT,
    experience TEXT,
    education TEXT,

    linkedin_url TEXT,
    github_url TEXT,
    portfolio_url TEXT,

    profile_image_url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_job_seeker_profiles_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
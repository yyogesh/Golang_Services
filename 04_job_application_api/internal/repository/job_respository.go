package repository

import (
	"database/sql"
	"job_portal/internal/models"
)

func CreateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	result, err := db.Exec("INSERT INTO jobs (title, description, company, location, min_salary, max_salary, type, experience_level, skills, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", job.Title,
		job.Description, job.Company, job.Location, job.MinSalary, job.MaxSalary, job.Type, job.ExperienceLevel, job.Skills, job.UserID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	job.ID = int(id)
	return job, nil
}

func GetAllJobs(db *sql.DB) ([]models.Job, error) {
	rows, err := db.Query("SELECT * FROM jobs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Company, &job.Location, &job.MinSalary, &job.MaxSalary, &job.Type, &job.ExperienceLevel, &job.Skills, &job.UserID, &job.CreatedAt); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func GetAllJobsByUserID(db *sql.DB, userID int) ([]models.Job, error) {
	rows, err := db.Query("SELECT * FROM jobs WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.MinSalary, &job.MaxSalary, &job.Type, &job.ExperienceLevel, &job.Skills, &job.UserID, &job.CreatedAt); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func GetJobByID(db *sql.DB, id int) (*models.Job, error) {
	job := &models.Job{}
	row := db.QueryRow("SELECT * FROM jobs WHERE id = ?", id)

	if err := row.Scan(&job.ID, &job.Title, &job.Description, &job.Location, &job.Company, &job.MinSalary, &job.MaxSalary, &job.Type, &job.ExperienceLevel, &job.Skills, &job.UserID, &job.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return job, nil
}

func UpdateJob(db *sql.DB, job *models.Job) (*models.Job, error) {
	_, err := db.Exec("UPDATE jobs SET title = ?, description = ?, company = ?, location = ?, min_salary = ?, max_salary = ?, type = ?, experience_level = ?, skills = ? WHERE id = ?", job.Title,
		job.Description, job.Company, job.Location, job.MinSalary, job.MaxSalary, job.Type, job.ExperienceLevel, job.Skills, job.ID)

	if err != nil {
		return nil, err
	}

	return job, nil
}

func DeleteJob(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM jobs WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

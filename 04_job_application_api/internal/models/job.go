package models

import "time"

type Job struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	Company         string    `json:"company"`
	MinSalary       int       `json:"min_salary"`
	ExperienceLevel string    `json:"experience_level"`
	Skills          string    `json:"skills"`
	MaxSalary       int       `json:"max_salary"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
	UserID          int       `json:"user_id"`
}

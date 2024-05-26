package models

import (
	"time"

	"gorm.io/gorm"
)

type GitQuery struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

type GitResponse struct {
	Data struct {
		User struct {
			Login                   string `json:"login"`
			Typename                string `json:"__typename"`
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
			Followers struct {
				TotalCount int `json:"totalCount"`
			} `json:"followers"`
			Repositories struct {
				Nodes []struct {
					Name           string `json:"name"`
					StargazerCount int    `json:"stargazerCount"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// User represents a user in the system (recruiter or candidate)
type User struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	Username    string `gorm:"size:50;not null;unique"`
	Password    string `gorm:"size:100;not null"` // This should be a hashed password
	CompanyName string `gorm:"size:100"`
}

// AuthToken represents an authentication token for a user
type AuthToken struct {
	gorm.Model
	UserID    uint   `gorm:"not null;index"`
	AuthToken string `gorm:"size:255;not null;unique"`
	ExpiresAt time.Time
}

// PasswordReset represents a password reset token for a user
type PasswordReset struct {
	gorm.Model
	UserID     uint   `gorm:"not null;index"`
	ResetToken string `gorm:"size:255;not null;unique"`
	ExpiresAt  time.Time
}

// Job represents a job posting by a recruiter
type Job struct {
	gorm.Model
	Title         string         `gorm:"size:255;not null"`
	Description   string         `gorm:"type:text;not null"`
	IsActive      bool           `gorm:"default:true"`
	DriveLink     string         `gorm:"size:255;not null"`
	UserID        uint           `gorm:"not null;index"`
	TopCandidates []TopCandidate // This indicates the one-to-many relationship
}

// TopCandidate represents top candidates for a job
type TopCandidate struct {
	gorm.Model
	JobID              uint   `gorm:"index"`
	Name               string `gorm:"size:100;not null"`
	GithubLink         string `gorm:"size:255;not null"`
	TotalContributions int
	Followers          int
	BestRepository     string `gorm:"size:255"`
	TotalStars         int
}

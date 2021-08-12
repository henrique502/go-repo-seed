package domain

import "time"

type Alert struct {
	ID            string
	Name          string
	Priority      string
	Source        string
	Message       string
	IntegrationID string
	ResponderIDs  []string
	ColletedAt    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

package domain

import "time"

type Alert struct {
	Id            string
	Name          string
	Priority      string
	Source        string
	Message       string
	IntegrationId string
	ResponderIds  []string
	ColletedAt    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

package sync

import "time"

type SyncApp struct{}

type Sync interface {
	AlertsByDay(day time.Time) error
	Alerts() error
	Integrations() error
	Teams() error
}

func New() SyncApp {
	return SyncApp{}
}

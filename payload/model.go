package payload

import (
	"time"
)

type Project struct {
	Name       string
	Code       string
	Building   bool
	Type       string
	Repository string
	Statuses   []ProjectStatus
}

type ProjectStatus struct {
	// ProjectId   int
	BuildId     int
	Success     bool
	Url         string
	PublishedAt time.Time
}

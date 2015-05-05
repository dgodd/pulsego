package payload

import "time"

type ProjectStatus struct {
	Success     bool
	Url         string
	PublishedAt time.Time
	ProjectId   int
	BuildId     int
}

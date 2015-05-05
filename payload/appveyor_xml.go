package payload

import (
	"encoding/json"
	"time"
)

type build struct {
	BuildId     int    `json:"buildId"`
	BuildNumber int    `json:"buildNumber"`
	Branch      string `json:"branch"`
	Status      string `json:"status"`
	Created     string `json:"created"`
}

func AppVeyorXmlPayload(xmlContent []byte) ([]ProjectStatus, error) {
	var a struct {
		Builds []build `json:"builds"`
	}
	err := json.Unmarshal(xmlContent, &a)
	if err != nil {
		return nil, err
	}
	longform := "2006-01-02T15:04:05.0000000-07:00"
	statuses := make([]ProjectStatus, len(a.Builds))
	for i := 0; i < len(a.Builds); i++ {
		// statuses[i].ProjectId = ???
		statuses[i].BuildId = a.Builds[i].BuildId
		publishedAt, err := time.Parse(longform, a.Builds[i].Created)
		if err == nil {
			statuses[i].PublishedAt = publishedAt
		}
		statuses[i].Success = a.Builds[i].Status == "success"
	}
	return statuses, nil
}

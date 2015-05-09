package download

import (
	"io/ioutil"
	"net/http"

	"github.com/dgodd/pulsego/payload"
)

func Download(project payload.Project) ([]payload.ProjectStatus, error) {
	resp, err := http.Get("https://ci.appveyor.com/api/projects/" + project.Repository + "/history?recordsNumber=10")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	statuses, err := payload.AppVeyorXmlPayload(body)
	if err != nil {
		panic(err)
	}
	return statuses, err
}

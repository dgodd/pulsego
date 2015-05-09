package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dgodd/pulsego/download"
	"github.com/dgodd/pulsego/payload"
	"github.com/unrolled/render"
)

func Downloads(projects []*payload.Project) {
	for _, project := range projects {
		statuses, isBuilding, err := download.Download(*project)
		if err != nil {
			panic(err)
		}
		project.Building = isBuilding
		project.Statuses = statuses
	}
}

func main() {
	projects := make([]*payload.Project, 0)
	projects = append(projects, &payload.Project{
		Name:       "Pulse",
		Code:       "PLSE",
		Type:       "appveyor",
		Repository: "dgodd/pulsego",
	})
	go Downloads(projects)

	r := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/api/projects", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, projects)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type resource struct {
	ID      int     `json:"id"`
	Name string `json:"name"`
	FilePath string `json:"file_path"`
	Project int `json:"project"`
}

type resourceList []resource

var resources = resourceList{
	{
		ID: 1,
		Name: "UML ER",
		FilePath: "/res/12dasds4",
		Project: 1,
	},{
		ID: 2,
		Name: "Design",
		FilePath: "/res/12dasds4dsa",
		Project: 1,
	},{
		ID: 3,
		Name: "UML ER",
		FilePath: "/res/12asddasds4",
		Project: 3,
	},{
		ID: 4,
		Name: "UML ER",
		FilePath: "/res/12asddasds4",
		Project: 4,
	},
}

var resourcesIndexer = 4

func HandleResourceCreate(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	isAuthenticated, user := AuthoriseByToken(r)

	if isAuthenticated {
		var newResource resource
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, &newResource)

		if newResource.Name == "" || newResource.FilePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, p := range projects {
			if p.ID == newResource.Project {
				if p.Owner == user.ID {
					w.Header().Add("Content-Type", "application/json")
					resources = append(resources, newResource)
					newResource.ID = resourcesIndexer
					resourcesIndexer++
					w.WriteHeader(http.StatusCreated)
					return
				}
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
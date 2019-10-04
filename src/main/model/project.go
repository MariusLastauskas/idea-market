package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type project struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	IsPublic bool `json:"is_public"`
	Price float32 `json:"price"`
	Multiplicity int `json:"multiplicity"`
	Owner int `json:"owner"`

	Buyers []int `json:"buyers"`
}

type projectList []project

var projects = projectList{
	{
		ID: 1,
		Name: "Awesome project with go for unlimited sale",
		Description: "It involves e-shopping",
		IsPublic: true,
		Price: 19.99,
		Multiplicity: 0,
		Owner: 1,
		Buyers: []int{2},
	},
	{
		ID: 2,
		Name: "Even better project with limited edition",
		Description: "It involves AI",
		IsPublic: true,
		Price: 99.99,
		Multiplicity: 5,
		Owner: 1,
		Buyers: []int{2, 3},
	},
	{
		ID: 3,
		Name: "Even better project with limited edition, sold out",
		Description: "It involves AI",
		IsPublic: true,
		Price: 99.99,
		Multiplicity: 2,
		Owner: 4,
		Buyers: []int{2, 3},
	},
	{
		ID: 4,
		Name: "Even better project free",
		Description: "It involves AI",
		IsPublic: true,
		Price: 0,
		Owner: 4,
		Multiplicity: 0,
	},
	{
		ID: 5,
		Name: "Even better project with limited edition",
		Description: "It involves AI",
		IsPublic: false,
		Price: 99.99,
		Multiplicity: 5,
		Owner: 1,
		Buyers: []int{2, 3},
	},
	{
		ID: 6,
		Name: "Even better project with limited edition",
		Description: "It involves AI",
		IsPublic: false,
		Price: 99.99,
		Multiplicity: 5,
		Owner: 2,
		Buyers: []int{2, 3},
	},
}

var projectIndexer = 7

func HandleProjectsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		resultProjects := projectList{}

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		for _, p := range projects {
			isAuthorised, u := AuthoriseByToken(r)
			if isAuthorised && u.Role == 1 || isAuthorised && p.Owner == u.ID || p.IsPublic {
				resultProjects = append(resultProjects, p)
			}
		}

		json.NewEncoder(w).Encode(resultProjects)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleProjectCreate(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	isAuthenticated, user := AuthoriseByToken(r)

	if isAuthenticated {
		var newProject project
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, &newProject)

		if newProject.Name == "" || newProject.Description == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newProject.ID = projectIndexer
		newProject.Owner = user.ID
		projectIndexer++
		projects = append(projects, newProject)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newProject)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func HandleProjectRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	subpath, isSubpath := getSubroute(r.RequestURI)

	switch r.Method {
	case "GET":
		if isSubpath {
			id, err := getSubIdFromSubpath(subpath)
			if subpath == "donations" {
				d, status := getProjectDonations(r, -1)
				w.WriteHeader(status)

				if status == http.StatusOK {
					json.NewEncoder(w).Encode(d)
				}
			} else if err == nil && id > -1 {
				d, status := getProjectDonations(r, id)
				w.WriteHeader(status)

				if status == http.StatusOK {
					json.NewEncoder(w).Encode(d)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			p, status := projectGet(r)
			w.WriteHeader(status)

			if status == http.StatusOK {
				json.NewEncoder(w).Encode(p)
			}
		}
	case "DELETE":
		p, status := projectDelete(r)
		w.WriteHeader(status)

		if status == http.StatusNoContent {
			json.NewEncoder(w).Encode(p)
		}
	case "PUT":
		p, status := projectUpdate(r)
		w.WriteHeader(status)

		if status == http.StatusOK {
			json.NewEncoder(w).Encode(p)
		}
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func projectGet(r *http.Request) (project, int) {
	id, err := GetIdFromUrl(r.RequestURI)
	isAuthenticated, user := AuthoriseByToken(r)

	if err != nil || id == -1 {
		return project{}, http.StatusBadRequest
	}

	for _, p := range projects {
		if p.ID == id {
			if p.IsPublic || isAuthenticated && p.Owner == user.ID || isAuthenticated && user.Role == 1 {
				return p, http.StatusOK
			}

			return project{}, http.StatusForbidden
		}
	}

	return project{}, http.StatusNotFound
}

func getProjectDonations(r *http.Request, d_id int) (donationsList, int) {
	p_id, err := GetIdFromUrl(r.RequestURI)
	isAuthenticated, user := AuthoriseByToken(r)
	resultDonations := donationsList{}
	parts := strings.Split(r.RequestURI, "/")

	if err != nil || p_id == -1 {
		return donationsList{}, http.StatusBadRequest
	}

	if parts[3] != "donations" && parts[3] != "donation" || parts[3] != "donations" && d_id < 0 || parts[3] != "donation" && d_id > -1 {
		return donationsList{}, http.StatusNotFound
	}

	for _, p := range projects {
		if p.ID == p_id {
			if p.IsPublic || isAuthenticated && p.Owner == user.ID || isAuthenticated && user.Role == 1 {
				for _, d :=range donations {
					if d.Project == p_id && d_id == -1 || d.Project == p_id && d_id == d.ID{
						resultDonations = append(resultDonations, d)
					}
				}
				if len(resultDonations) == 0 {
					return donationsList{}, http.StatusNotFound
				}
				return resultDonations, http.StatusOK
			}

			return donationsList{}, http.StatusForbidden
		}
	}
	return donationsList{}, http.StatusNotFound
}

func projectDelete(r *http.Request) (project, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return project{}, http.StatusBadRequest
	}

	isAuthorised, u := AuthoriseByToken(r)
	for i, p := range projects {
		if p.ID == id {
			if isAuthorised && u.Role == 1 || isAuthorised && u.ID == p.Owner {
				projects = append(projects[:i], projects[i+1:]...)
				return p, http.StatusNoContent
			}
			return project{}, http.StatusForbidden
		}
	}

	return project{}, http.StatusNotFound
}

func projectUpdate(r *http.Request) (project, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return project{}, http.StatusBadRequest
	}

	isAuthorised, u := AuthoriseByToken(r)
	var updateProjects project
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return project{}, http.StatusBadRequest
	}
	json.Unmarshal(reqBody, &updateProjects)

	if updateProjects.Name != "" && updateProjects.Description != "" {
		for i, p := range projects {
			if p.ID == id {
				if isAuthorised && u.Role == 1 || isAuthorised && u.ID == p.Owner {
					p.Name = updateProjects.Name
					p.Description = updateProjects.Description
					p.IsPublic = updateProjects.IsPublic
					p.Multiplicity = updateProjects.Multiplicity
					p.Price = updateProjects.Price
					p.Owner = updateProjects.Owner
					p.Buyers = updateProjects.Buyers

					remProjects := projects[i+1:]
					projects = append(projects[:i], p)
					projects = append(projects, remProjects...)
					return p, http.StatusOK
				}

				return project{}, http.StatusForbidden
			}
		}
		return project{}, http.StatusNotFound
	}

	return project{}, http.StatusBadRequest
}
package model

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type resource struct {
	ID      int     `json:"id"`
	Name string `json:"name"`
	FilePath string `json:"file_path"`
	Project int `json:"project"`
}

type resourceList []resource

func HandleResourceCreate(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("access-control-allow-origin", "http://localhost:3000")
	w.Header().Set("access-control-allow-methods", "GET, OPTIONS, POST, PATCH, PUT, DELETE");
	w.Header().Set("access-control-allow-headers", "Origin, Content-Type, X-Auth-Token");
	w.Header().Set("access-control-allow-credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK);
		return
	}
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

		p := getProjectsList("select * from project where id_project=" + strconv.Itoa(newResource.Project))

		if len(p) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if p[0].Owner.ID == user.ID {
			w.Header().Add("Content-Type", "application/json")

			db, err := sql.Open("mysql", "root:@/saitynai")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			resourceIdRows, err := db.Query("SELECT id_Resource from resource ORDER BY id_Resource DESC LIMIT 1")
			if err != nil {
				log.Fatal(err)
			}

			for resourceIdRows.Next() {
				err = resourceIdRows.Scan(&newResource.ID)
				if err != nil {
					log.Fatal(err)
				}
				newResource.ID++

				db.Query("INSERT INTO resource (`name`, file_path, id_Resource, fk_project) VALUES (?, ?, ?, ?)", newResource.Name, newResource.FilePath, newResource.ID, newResource.Project)

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)

				json.NewEncoder(w).Encode(newResource)
			}

			return
		}
		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func getResourcesList(s string) resourceList {
	var (
		id_Resource int
		name string
		file_path string
		fk_project int
	)

	db, err := sql.Open("mysql", "root:@/saitynai")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var resources = resourceList{}
	var r = resource{}
	for rows.Next() {
		err := rows.Scan(&name, &file_path, &id_Resource, &fk_project)
		if err != nil {
			log.Fatal(err)
		}
		r = resource{
			ID: id_Resource,
			Name: name,
			FilePath: file_path,
			Project: fk_project,
		}
		resources = append(resources, r)
	}
	return resources
}
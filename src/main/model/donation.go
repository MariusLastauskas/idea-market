package model

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type donation struct {
	ID      int     `json:"id"`
	Size    float32 `json:"size"`
	Donor int     `json:"donator"`
	Project int     `json:"project"`
}

type donationsList []donation

func HandleDonationsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(getDonationsList("select * from donation"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleDonationCreate(w http.ResponseWriter, r *http.Request)  {
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
		var newDonation donation
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, &newDonation)
		newDonation.Donor = user.ID

		if newDonation.Size <= 0 || newDonation.Project < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p := getProjectsList("select * from project where id_project=" + strconv.Itoa(newDonation.Project))

		if len(p) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		db, err := sql.Open("mysql", "root:@/saitynai")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		donationIdRow, err := db.Query("SELECT id_Donation from donation ORDER BY id_donation DESC LIMIT 1")
		if err != nil {
			log.Fatal(err)
		}

		for donationIdRow.Next() {
			err = donationIdRow.Scan(&newDonation.ID)
			if err != nil {
				log.Fatal(err)
			}
			newDonation.ID++

			db.Query("insert into donation (size, id_donation, fk_donor, fk_project) values (?, ?, ?, ?)", newDonation.Size,newDonation.ID,newDonation.Donor, newDonation.Project)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(newDonation)
		}
		return
	}
	w.WriteHeader(http.StatusForbidden)
}

func getDonationsList(s string) donationsList {
	var (
		size float32
		id_donation int
		fk_donor int
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

	var donations = donationsList{}
	var d = donation{}
	for rows.Next() {
		err := rows.Scan(&size, &id_donation, &fk_donor, &fk_project)
		if err != nil {
			log.Fatal(err)
		}
		d = donation{
			ID: id_donation,
			Size: size,
			Donor: fk_donor,
			Project: fk_project,
		}
		donations = append(donations, d)
	}
	return donations
}
package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type donation struct {
	ID      int     `json:"id"`
	Size    float32 `json:"size"`
	Donor int     `json:"donator"`
	Project int     `json:"project"`
}

type donationsList []donation

var donations = donationsList{
	{
		ID: 1,
		Size: 3.00,
		Donor: 1,
		Project: 2,
	},
	{
		ID: 2,
		Size: 3.00,
		Donor: 2,
		Project: 2,
	},
	{
		ID: 3,
		Size: 3.00,
		Donor: 1,
		Project: 1,
	},
	{
		ID: 4,
		Size: 3.00,
		Donor: 1,
		Project: 1,
	},
}

var donationsIndexer = 5

func HandleDonationsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(donations)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleDonationCreate(w http.ResponseWriter, r *http.Request)  {
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

		for _, p := range projects {
			if p.ID == newDonation.Project {
				w.Header().Add("Content-Type", "application/json")
				donations = append(donations, newDonation)
				newDonation.ID = donationsIndexer
				projectIndexer++
				w.WriteHeader(http.StatusCreated)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
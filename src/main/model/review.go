package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type review struct {
	ID int `json:"id"`
	Comment string `json:"comment"`
	Grade int `json:"grade"`
	Reviewer int `json:"reviewer"`
	Project int `json:"project"`
}

type reviewList []review

var reviews = reviewList{
	{
		ID:      1,
		Comment: "Me likey",
		Grade: 5,
		Reviewer: 1,
		Project: 4,
	},
	{
		ID: 2,
		Comment: "Me no likey",
		Grade: 1,
		Reviewer: 2,
		Project: 4,
	},
	{
		ID: 3,
		Comment: "Me somewhat likey",
		Grade: 3,
		Reviewer: 3,
		Project: 5,
	},
}

var reviewIndexer = 4

func HandleReviewCreate(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	isAuthenticated, user := AuthoriseByToken(r)

	if isAuthenticated {
		var newReview review
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, &newReview)
		newReview.Reviewer = user.ID

		//for _, p := range projects {
		//	if p.ID == newReview.Project {
		//		w.Header().Add("Content-Type", "application/json")
		//		newReview.ID = reviewIndexer
		//		reviewIndexer++
		//		reviews = append(reviews, newReview)
		//		w.WriteHeader(http.StatusCreated)
		//		json.NewEncoder(w).Encode(newReview);
		//		return
		//	}
		//}
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
package model

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type review struct {
	ID int `json:"id"`
	Comment string `json:"comment"`
	Grade int `json:"grade"`
	Reviewer int `json:"reviewer"`
	Project int `json:"project"`
}

type reviewList []review

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

		p := getProjectsList("select * from project where id_project=" + strconv.Itoa(newReview.Project))

		if len(p) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if user.IsActive {
			w.Header().Add("Content-Type", "application/json")

			db, err := sql.Open("mysql", "root:@/saitynai")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			reviewIdRows, err := db.Query("SELECT id_Review from review ORDER BY id_Review DESC LIMIT 1")
			if err != nil {
				log.Fatal(err)
			}

			for reviewIdRows.Next() {
				err = reviewIdRows.Scan(&newReview.ID)
				if err != nil {
					log.Fatal(err)
				}
				newReview.ID++

				db.Query("INSERT INTO review (comment, grade, id_review, fk_user, fk_project) values (?, ?, ?, ?, ?)", newReview.Comment, newReview.Grade, newReview.ID, user.ID, p[0].ID)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)

				json.NewEncoder(w).Encode(newReview)
			}

			return
		}

		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func getReviewsList(s string) reviewList {
	var (
		comment string
		grade int
		id_review int
		fk_user int
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

	var reviews = reviewList{}
	var r = review{}
	for rows.Next() {
		err := rows.Scan(&comment, &grade, &id_review, &fk_user, &fk_project)
		if err != nil {
			log.Fatal(err)
		}
		r = review{
			ID:id_review,
			Comment: comment,
			Grade: grade,
			Reviewer: fk_user,
			Project: fk_project,
		}
		reviews = append(reviews, r)
	}
	return reviews
}
package main

import (
	"./model"
	"log"
	"net/http"
)

func main() {
	// 404
	http.HandleFunc("/", model.HandleBadPath)

	// login
	http.HandleFunc("/login", model.HandleLogin)

	// articles
	http.HandleFunc("/articles", model.HandleArticlesGet)
	http.HandleFunc("/article", model.HandleArticleCreate)
	http.HandleFunc("/article/", model.HandleArticleRequest)

	// users
	http.HandleFunc("/users", model.HandleUsersGet)
	http.HandleFunc("/user", model.HandleUserCreate)
	http.HandleFunc("/user/", model.HandleUserRequest)

	// projects
	http.HandleFunc("/projects", model.HandleProjectsGet)
	http.HandleFunc("/project", model.HandleProjectCreate)
	http.HandleFunc("/project/", model.HandleProjectRequest)

	// donations
	http.HandleFunc("/donation", model.HandleDonationCreate)

	// review
	http.HandleFunc("/review", model.HandleReviewCreate)

	// resources
	http.HandleFunc("/resource", model.HandleResourceCreate)

	http.Header{}.Set("Access-Control-Allow-Origin", "*")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
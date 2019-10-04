package main

import (
	"./model"
	"log"
	"net/http"
)

func main() {
	// 404
	http.HandleFunc("/", model.HandleBadPath)

	// articles
	http.HandleFunc("/articles", model.HandleArticlesGet)
	http.HandleFunc("/article", model.HandleArticleCreate)
	http.HandleFunc("/article/", model.HandleArticleRequest)

	// users
	http.HandleFunc("/users", model.HandleUsersGet)
	http.HandleFunc("/user", model.HandleUserCreate)
	http.HandleFunc("/user/", model.HandleUserRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
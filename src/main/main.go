package main

import (
	"./model"
	"log"
	"net/http"
)















//func handleArticle(w http.ResponseWriter, r *http.Request)  {
//	var x = r.URL.Path
//	print(x);
//	switch r.Method {
//	case "GET":
//
//	case "POST":
//
//	case "PUT":
//
//	default:
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}

func main() {
	// 404
	http.HandleFunc("/", model.HandleBadPath)

	// articles
	http.HandleFunc("/articles", model.HandleArticlesGet)
	http.HandleFunc("/article", model.HandleArticleCreate)
	http.HandleFunc("/article/", model.HandleArticleRequest)

	// users
	http.HandleFunc("/users", model.HandleUsersGet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
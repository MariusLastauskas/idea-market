package model

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	FullText string `json:"full_text"`
	IsPublic bool `json:"is_public"`
	Author int `json:"author"`
}

type articleList []article

func authoriseArticleBehaviour(r *http.Request, id int) (bool, User) {
	a := getArticlesList("select * from article where id_article=" + strconv.Itoa(id))

	if len(a) == 0 {
		return false, User{}
	}

	if a[0].ID == id && a[0].IsPublic {
		return true, User{}
	}

	isAuthenticated, u := AuthoriseByToken(r)
	if isAuthenticated && u.Role == 1 || isAuthenticated && u.ID == a[0].Author {
		return true, u
	}

	return false, User{}
}

func HandleArticlesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == "GET" {
		resultArticles := articleList{}

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		articles := getArticlesList("select * from article")

		for _, a := range articles {
			isAuthorised, _ := authoriseArticleBehaviour(r, a.ID)
			if isAuthorised {
				resultArticles = append(resultArticles, a)
			}
		}

		json.NewEncoder(w).Encode(resultArticles)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleArticleCreate(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	isAuthenticated, user := AuthoriseByToken(r)

	if isAuthenticated && user.Role == 1 {
		var newArticle article
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, &newArticle)

		if newArticle.Title == "" || newArticle.Content == "" || newArticle.FullText == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		db, err := sql.Open("mysql", "root:@/saitynai")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		articleIdRows, err := db.Query("SELECT id_article from article ORDER BY id_article DESC LIMIT 1")
		if err != nil {
			log.Fatal(err)
		}

		for articleIdRows.Next() {
			err = articleIdRows.Scan(&newArticle.ID)
			if err != nil {
				log.Fatal(err)
			}
			newArticle.ID++
		}

		newArticle.Author = user.ID

		db.Query("insert into article (title, content, full_text, is_public, id_article, fk_author) values (?, ?, ?, ?, ?, ?)", newArticle.Title, newArticle.Content, newArticle.FullText, newArticle.IsPublic, newArticle.ID, newArticle.Author)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newArticle)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func HandleArticleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		a, status := articleGet(r)
		w.WriteHeader(status)

		if status == http.StatusOK {
			json.NewEncoder(w).Encode(a)
		}
	case "DELETE":
		a, status := articleDelete(r)
		w.WriteHeader(status)

		if status == http.StatusNoContent {
			json.NewEncoder(w).Encode(a)
		}
	case "PUT":
		a, status := articleUpdate(r)
		w.WriteHeader(status)

		if status == http.StatusOK {
			json.NewEncoder(w).Encode(a)
		}
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func articleGet(r *http.Request) (article, int) {
	id, err := GetIdFromUrl(r.RequestURI)
	isAuthorised, user := AuthoriseByToken(r)

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	articles := getArticlesList("select * from article where id_article=" + strconv.Itoa(id))

	if len(articles) == 0 {
		return article{}, http.StatusNotFound
	}

	a := articles[0]

	if a.IsPublic || isAuthorised && a.Author == user.ID || isAuthorised && user.Role == 1 {
		return a, http.StatusOK
	}

	return article{}, http.StatusForbidden
}

func articleDelete(r *http.Request) (article, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	isAuthorised, u := AuthoriseByToken(r)

	db, err := sql.Open("mysql", "root:@/saitynai")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	articleToDelete := getArticlesList("select * from article where id_article=" + strconv.Itoa(id))

	if len(articleToDelete) == 0 {
		return article{}, http.StatusNotFound
	}

	if isAuthorised && u.Role == 1 || isAuthorised && u.ID == articleToDelete[0].Author {
		db.Query("delete from article where id_article=?", id)
		return article{}, http.StatusNoContent
	}

	return article{}, http.StatusForbidden
}

func articleUpdate(r *http.Request) (article, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	isAuthorised, u := AuthoriseByToken(r)
	if isAuthorised && u.Role == 1 {
		var updateArticle article
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return article{}, http.StatusBadRequest
		}
		json.Unmarshal(reqBody, &updateArticle)

		db, err := sql.Open("mysql", "root:@/saitynai")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		articleId, err := db.Query("select id_article, fk_author from article where id_article=?", id)
		if err != nil {
			log.Fatal(err)
		}

		for articleId.Next() {
			var aId int
			var aAuthorId int
			articleId.Scan(&aId, &aAuthorId)
			if isAuthorised && u.Role == 1 {
				 _, err := db.Query("update article set title=?, content=?, full_text=?, is_public=?, fk_author=? " +
					"where id_article=?", updateArticle.Title, updateArticle.Content, updateArticle.FullText, updateArticle.IsPublic, updateArticle.Author,
					id)
				 if err != nil {
					log.Fatal(err)
				 }
				 return getArticlesList("select * from article where id_article=" + strconv.Itoa(id))[0], http.StatusOK
			}

			return article{}, http.StatusForbidden
		}

		return article{}, http.StatusNotFound
	}

	return article{}, http.StatusForbidden
}

func getArticlesList(s string) articleList {
	var (
		id_article int
		title string
		content string
		full_text string
		is_public bool
		fk_author int
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

	var articles = articleList{}
	var a = article{}
	for rows.Next() {
		err := rows.Scan(&title, &content, &full_text, &is_public, &id_article, &fk_author)
		if err != nil {
			log.Fatal(err)
		}
		a = article{
			ID: id_article,
			Title: title,
			Content: content,
			FullText: full_text,
			IsPublic: is_public,
			Author: fk_author,
		}

		articles = append(articles, a)
	}
	return articles
}
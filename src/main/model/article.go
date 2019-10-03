package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	FullText string `json:"full_text"`
	IsPublic bool `json:"is_public"`
}

type articleList []article

var articles = articleList{
	{
		ID: 1,
		Title: "Article about Animals",
		Content: "Cats and Dogs are #1 inheritance ref",
		FullText: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
			"labore et dolore magna aliqua. Diam donec adipiscing tristique risus nec feugiat. Maecenas accumsan " +
			"lacus vel facilisis volutpat est velit egestas. Libero volutpat sed cras ornare arcu. Morbi tristique " +
			"senectus et netus et malesuada fames. In metus vulputate eu scelerisque. Pretium lectus quam id leo in " +
			"vitae. Varius quam quisque id diam vel. At in tellus integer feugiat scelerisque varius morbi enim. " +
			"Augue neque gravida in fermentum et sollicitudin ac orci. Viverra orci sagittis eu volutpat odio " +
			"facilisis. Auctor augue mauris augue neque. Nisl tincidunt eget nullam non nisi est sit amet facilisis. " +
			"Posuere ac ut consequat semper viverra nam. Lacus laoreet non curabitur gravida arcu ac tortor " +
			"dignissim. Purus viverra accumsan in nisl nisi scelerisque.",
		IsPublic: true,
	},
	{
		ID: 2,
		Title: "Article about passwords",
		Content: "Treat them as your underwear",
		FullText: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
			"labore et dolore magna aliqua. Diam donec adipiscing tristique risus nec feugiat. Maecenas accumsan " +
			"lacus vel facilisis volutpat est velit egestas. Libero volutpat sed cras ornare arcu. Morbi tristique " +
			"senectus et netus et malesuada fames. In metus vulputate eu scelerisque. Pretium lectus quam id leo in " +
			"vitae. Varius quam quisque id diam vel. At in tellus integer feugiat scelerisque varius morbi enim. " +
			"Augue neque gravida in fermentum et sollicitudin ac orci. Viverra orci sagittis eu volutpat odio " +
			"facilisis. Auctor augue mauris augue neque. Nisl tincidunt eget nullam non nisi est sit amet facilisis. " +
			"Posuere ac ut consequat semper viverra nam. Lacus laoreet non curabitur gravida arcu ac tortor " +
			"dignissim. Purus viverra accumsan in nisl nisi scelerisque.",
		IsPublic: true,
	},
	{
		ID: 3,
		Title: "Unfinished article",
		Content: "Super secret content",
		FullText: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
			"labore et dolore magna aliqua. Diam donec adipiscing tristique risus nec feugiat. Maecenas accumsan " +
			"lacus vel facilisis volutpat est velit egestas. Libero volutpat sed cras ornare arcu. Morbi tristique " +
			"senectus et netus et malesuada fames. In metus vulputate eu scelerisque. Pretium lectus quam id leo in " +
			"vitae. Varius quam quisque id diam vel. At in tellus integer feugiat scelerisque varius morbi enim. " +
			"Augue neque gravida in fermentum et sollicitudin ac orci. Viverra orci sagittis eu volutpat odio " +
			"facilisis. Auctor augue mauris augue neque. Nisl tincidunt eget nullam non nisi est sit amet facilisis. " +
			"Posuere ac ut consequat semper viverra nam. Lacus laoreet non curabitur gravida arcu ac tortor " +
			"dignissim. Purus viverra accumsan in nisl nisi scelerisque.",
		IsPublic: false,
	},
}

var articlesIndexer = 4

func HandleArticlesGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(articles)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleArticleCreate(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newArticle article
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &newArticle)

	if newArticle.Content == "" || newArticle.Title == "" || newArticle.Content == "" || newArticle.FullText == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newArticle.ID = articlesIndexer
	articlesIndexer++
	articles = append(articles, newArticle)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newArticle)
}

func HandleArticleRequest(w http.ResponseWriter, r *http.Request) {
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

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	for _, a := range articles {
		if a.ID == id {
			return a, http.StatusOK
		}
	}

	return article{ID: -1}, http.StatusNotFound
}

func articleDelete(r *http.Request) (article, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	for i, a := range articles {
		if a.ID == id {
			articles = append(articles[:i], articles[i + 1:]...)
			return a, http.StatusNoContent
		}
	}

	return article{ID: -1}, http.StatusNotFound
}

func articleUpdate(r *http.Request) (article, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return article{}, http.StatusBadRequest
	}

	var updateArticle article
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return article{}, http.StatusNotModified
	}
	json.Unmarshal(reqBody, &updateArticle)

	for i, a := range articles {
		if a.ID == id {
			a.Title = updateArticle.Title
			a.Content = updateArticle.Content
			remArticles := articles[i + 1:]
			articles = append(articles[:i], a)
			articles = append(articles, remArticles...)
			return a, http.StatusOK
		}
	}

	return article{ID: -1}, http.StatusNotFound
}
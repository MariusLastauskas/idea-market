package model

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIdFromUrl(url string) (int, error)  {
	parts := strings.Split(url, "/")
	if parts[2] == "" {
		return -1, nil
	}
	return strconv.Atoi(parts[2])
}

func AuthoriseByPassHash(passHash string) (string, bool) {
	return "54ds68a--authToken--4894dsa2", passHash == "sd54648sad1--randomHash--848d2s18asd1"
}

func AuthoriseByToken(r *http.Request) (bool, int)  {
	token, err := r.Cookie("AuthToken")

	if err != nil {
		return false, 0
	}

	if (token.Value == "54ds68a--userAuthToken--4894dsa2") {
		return true, 0
	}

	if (token.Value == "54ds68a--adminAuthToken--4894dsa2") {
		return true, 1
	}

	return false, 0
}

func HandleBadPath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
package model

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIdFromUrl(url string) (int, error) {
	parts := strings.Split(url, "/")
	if parts[2] == "" {
		return -1, nil
	}
	return strconv.Atoi(parts[2])
}

func getSubIdFromSubpath(url string) (int, error) {
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return -1, nil
	}
	if parts[1] == "" {
		return -1, nil
	}
	return strconv.Atoi(parts[1])
}

func getSubroute(url string) (string, bool) {
	parts := strings.Split(url, "/")
	if len(parts) <= 3 {
		return "", false
	}

	if len(parts) == 4 {
		return parts[3], true
	}

	return  parts[3] + "/" + parts[4], true
}

func AuthoriseByPassHash(passHash string) (string, bool) {
	return "54ds68a--authToken--4894dsa2", passHash == "sd54648sad1--randomHash--848d2s18asd1"
}

func AuthoriseByToken(r *http.Request) (bool, user)  {
	token, err := r.Cookie("AuthToken")
	id := -1

	if err != nil {
		return false, user{}
	}

	if (token.Value == "54ds68a--userAuthToken--4894dsa2") {
		id = 1
	}

	if (token.Value == "54ds68a--blockedUserAuthToken--4894dsa2") {
		id = 2
	}

	if (token.Value == "54ds68a--adminAuthToken--4894dsa2") {
		id = 3
	}

	if id != -1 {
		for _, u := range users {
			if u.ID == id {
				return true, u
			}
		}
	}

	return false, user{}
}

func HandleBadPath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
package model

import (
	"database/sql"
	"encoding/json"
	"github.com/gbrlsnchs/jwt/v3"
	"log"
	"net/http"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"encoding/base64"
)

type CustomPayload struct {
	jwt.Payload
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	PhotoPath string `json:"photo_path"`
	Role int `json:"role"`
	IsActive bool `json:"is_active"`
	ID int `json:"id"`
}

var hs = jwt.NewHS256([]byte("secret"))

func GetIdFromUrl(url string) (int, error) {
	parts := strings.Split(url, "/")
	if parts[2] == "" {
		return -1, nil
	}
	return strconv.Atoi(parts[2])
}

func getSubIdFromURI(url string) (int, error) {
	parts := strings.Split(url, "/")
	if len(parts) < 5 {
		return -1, nil
	}
	if parts[4] == "" {
		return -1, nil
	}
	return strconv.Atoi(parts[4])
}

func getSubroute(url string) (string, bool) {
	parts := strings.Split(url, "/")
	if len(parts) <= 3 {
		return "", false
	}

	return parts[3], true
}

func AuthoriseByPassHash(usr string, passHash string) (bool, json.Token) {
	var (
		id_User int
		full_name string
		username string
		email string
		password_hash string
		photo_path string
		role int
		is_active bool
	)

	db, err := sql.Open("mysql", "root:@/saitynai")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `user` WHERE username=? AND password_hash=?", usr, passHash)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&full_name, &username, &email, &password_hash, &photo_path, &role, &is_active, &id_User)
		if err != nil {
			log.Fatal(err)
		}

		now := time.Now()
		pl := CustomPayload{
			Payload: jwt.Payload{
				Issuer:         "idea market",
				Subject:        "user",
				Audience:       jwt.Audience{"http://localhost:8080/"},
				ExpirationTime: jwt.NumericDate(now.Add(1 * time.Hour)),
				NotBefore:      jwt.NumericDate(now.Add(1 * time.Hour)),
				IssuedAt:       jwt.NumericDate(now),
				JWTID:          usr,
			},
			FullName: full_name,
			Username: username,
			Email: email,
			PhotoPath: photo_path,
			Role: role,
			IsActive: is_active,
			ID: id_User,
		}

		token, err := jwt.Sign(pl, hs)
		if err == nil {
			return true, token
		}
	}

	return false, nil
}

func AuthoriseByToken(r *http.Request) (bool, User)  {
	var pl CustomPayload
	encodedToken, err := r.Cookie("jwtToken")
	if err != nil {
		return false, User{}
	}

	token, err := base64.StdEncoding.DecodeString(encodedToken.Value)
	hd, err := jwt.Verify(token, hs, &pl)
	hd.KeyID = "";

	if err != nil || pl.ExpirationTime.Time.UTC().Before(time.Now().UTC()) {
		return false, User{}
	}

	user := getUsersList("select * from User where id_User = " + strconv.Itoa(pl.ID))

	return true, user[0]
}

//func TokenToString(runes []byte) string {
//	outString := ""
//	for _, v := range runes {
//		outString += string(v)
//	}
//	return outString
//}

func HandleBadPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNotFound)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username, err_username := r.Cookie("username")
	passHash, err_pass_hash := r.Cookie("passHash")

	if err_username != nil || err_pass_hash != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	isAuthenticated, token := AuthoriseByPassHash(username.Value, passHash.Value)

	if isAuthenticated {
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(token)
		return
	}

	w.WriteHeader(http.StatusForbidden)
}

func contains(arr []int, el int) bool {
	for _, e := range arr {
		if e == el {
			return true
		}
	}
	return false
}
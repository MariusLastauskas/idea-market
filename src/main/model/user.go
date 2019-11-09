package model

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	ID int `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	PhotoPath string `json:"photo_path"`
	Role int `json:"role"`
	IsActive bool `json:"is_active"`

	Articles articleList `json:"articles"`
	OwnedProjects projectList `json:"owned_projects"`
	BoughtProjects projectList `json:"bought_projects"`
	FollowedProjects projectList `json:"followed_projects"`
}

type userList []user

func authoriseUserBehaviour(r *http.Request, id int) bool {
	isAuthenticated, u := AuthoriseByToken(r)
	return u.Role == 1 || isAuthenticated && u.ID == id && u.IsActive
}

func HandleUsersGet(w http.ResponseWriter, r *http.Request) {

 	isAuthenticated, u := AuthoriseByToken(r)

	if !isAuthenticated || u.Role != 1 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		users := getUsersList("select * from User")

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleUserCreate(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &newUser)

	if newUser.FullName == "" || newUser.Username == "" ||  newUser.Email == "" ||
		newUser.PasswordHash == "" || newUser.PhotoPath == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:@/saitynai")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userId, err := db.Query("SELECT id_User from `user` ORDER BY id_User DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	for userId.Next() {
		err = userId.Scan(&newUser.ID)
		if err != nil {
			log.Fatal(err)
		}
		newUser.ID++
	}

	newUser.Role = 0
	newUser.IsActive = true
	newUser.Articles = []article{}
	//userIndexer++

	db.Query("INSERT INTO `user` (`full_name`, `username`, `email`, `password_hash`, " +
		"`photo_path`, `role`, `is_active`, `id_User`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);", newUser.FullName,
		newUser.Username, newUser.Email, newUser.PasswordHash, newUser.PhotoPath, newUser.Role, newUser.IsActive,
		newUser.ID);

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
	return
}

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		u, status := userGet(r)
		w.WriteHeader(status)

		if status == http.StatusOK {
			json.NewEncoder(w).Encode(u)
		}
	case "DELETE":
		u, status := userDelete(r)
		w.WriteHeader(status)

		if status == http.StatusNoContent {
			json.NewEncoder(w).Encode(u)
		}
	case "PUT":
		u, status := userUpdate(r)
		w.WriteHeader(status)

		if status == http.StatusOK {
			json.NewEncoder(w).Encode(u)
		}
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func userGet(r *http.Request) (user, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return user{}, http.StatusBadRequest
	}

	if authoriseUserBehaviour(r, id) {
		users := getUsersList("select * from User where id_User = " + strconv.Itoa(id))
		if len(users) > 0 {
			return users[0], http.StatusOK
		}

		return user{}, http.StatusNotFound
	}

	return user{}, http.StatusForbidden
}

func userDelete(r *http.Request) (user, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return user{}, http.StatusBadRequest
	}

	if authoriseUserBehaviour(r, id) {
		db, err := sql.Open("mysql", "root:@/saitynai")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		userId, err := db.Query("SELECT id_User from `user` WHERE id_User=?", id)
		if err != nil {
			log.Fatal(err)
		}

		for userId.Next() {
			db.Query("DELETE FROM `user` WHERE `user`.`id_User` = ?", id)
			return user{}, http.StatusNoContent
		}

		return user{}, http.StatusNotFound
	}

	return user{}, http.StatusForbidden
}

func userUpdate(r *http.Request) (user, int) {
	id, err := GetIdFromUrl(r.RequestURI)

	if err != nil || id == -1 {
		return user{}, http.StatusBadRequest
	}

	if authoriseUserBehaviour(r, id) {
		var updateUser user
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return user{}, http.StatusNotModified
		}
		json.Unmarshal(reqBody, &updateUser)

		db, err := sql.Open("mysql", "root:@/saitynai")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		userId, err := db.Query("SELECT id_User from `user` WHERE id_User=?", id)
		if err != nil {
			log.Fatal(err)
		}

		for userId.Next() {
			db.Query("UPDATE `user` SET `full_name`=?, `username`=?, `email`=?, `password_hash`=?, " +
				"`photo_path`=?, `role`=?, `is_active`=? WHERE `user`.`id_User` = ?", updateUser.FullName,
				updateUser.Username, updateUser.Email, updateUser.PasswordHash, updateUser.PhotoPath, updateUser.Role,
				updateUser.IsActive, id);

			return getUsersList("select * from User where id_User = " + strconv.Itoa(id))[0], http.StatusOK
		}

		return user{}, http.StatusNotFound
	}

	return user{}, http.StatusForbidden
}

func getUsersList(s string) userList {
	var (
		id_User int
		full_name string
		username string
		email string
		password_hash string
		photo_path string
		role int
		is_active bool
		id_Article int
		title string
		content string
		full_text string
		is_public bool
		fk_Author int
		id_Project int
		name string
		description string
		price float32
		multiplicity int
		fk_Owner int
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

	var users = userList{}
	var u = user{}
	var a = article{}
	var op = project{}
	for rows.Next() {
		err := rows.Scan(&full_name, &username, &email, &password_hash, &photo_path, &role, &is_active, &id_User)
		if err != nil {
			log.Fatal(err)
		}
		u = user{
			ID: id_User,
			FullName: full_name,
			Username: username,
			Email: email,
			PasswordHash: password_hash,
			PhotoPath: photo_path,
			Role: role,
			IsActive: is_active,
			Articles: articleList{},
			OwnedProjects: projectList{},
			BoughtProjects: projectList{},
			FollowedProjects: projectList{},
		}

		articleRows, err := db.Query("select * from Article where fk_Author = ?", id_User);
		defer articleRows.Close()
		for articleRows.Next() {
			err = articleRows.Scan(&title, &content, &full_text, &is_public, &id_Article, &fk_Author)
			if err != nil {
				log.Fatal(err)
			}
			a = article{
				ID: id_Article,
				Title: title,
				Content: content,
				FullText: full_text,
				IsPublic: is_public,
			}
			u.Articles = append(u.Articles, a)
		}

		ownedProjectRows, _ := db.Query("select * from Project where fk_Owner = ?", id_User)
		defer ownedProjectRows.Close()
		for ownedProjectRows.Next() {
			err = ownedProjectRows.Scan(&name, &description, &is_public, &price, &multiplicity, &id_Project, &fk_Owner)
			if err != nil {
				log.Fatal(err)
			}
			op = project{
				ID: id_Project,
				Name: name,
				Description: description,
				IsPublic: is_public,
				Price: price,
				Multiplicity: multiplicity,
			}
			u.OwnedProjects = append(u.OwnedProjects, op)
		}

		boughtProjectRow, _ := db.Query(
			"SELECT name, description, is_public, price, multiplicity, id_Project, fk_Owner " +
				"FROM Project INNER JOIN bought_projects ON id_Project=fk_Project WHERE fk_Buyer=?",
			id_User)
		defer boughtProjectRow.Close()
		for boughtProjectRow.Next() {
			err = boughtProjectRow.Scan(&name, &description, &is_public, &price, &multiplicity, &id_Project, &fk_Owner)
			if err != nil {
				log.Fatal(err)
			}
			bp := project{
				ID: id_Project,
				Name: name,
				Description: description,
				IsPublic: is_public,
				Price: price,
				Multiplicity: multiplicity,
			}
			u.BoughtProjects = append(u.BoughtProjects, bp)
		}

		followedProjectRow, _ := db.Query(
			"SELECT name, description, is_public, price, multiplicity, id_Project, fk_Owner " +
				"FROM Project INNER JOIN folowed_projects ON id_Project=fk_Project WHERE fk_User=?",
			id_User)
		defer followedProjectRow.Close()
		for followedProjectRow.Next() {
			err = followedProjectRow.Scan(&name, &description, &is_public, &price, &multiplicity, &id_Project, &fk_Owner)
			if err != nil {
				log.Fatal(err)
			}
			fp := project{
				ID: id_Project,
				Name: name,
				Description: description,
				IsPublic: is_public,
				Price: price,
				Multiplicity: multiplicity,
			}
			u.FollowedProjects = append(u.FollowedProjects, fp)
		}

		users = append(users, u)
	}
	return users
}
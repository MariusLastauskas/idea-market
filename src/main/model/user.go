package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	BoughtProjects []int `json:"bought_projects"`
}

type userList []user

var users = userList{
	{
		ID: 1,
		FullName: "Vardenis Pavardenis",
		Username: "Var_denis",
		Email: "vardenis.pavardenis@gmail.com",
		PasswordHash: "65d64as8--userPassHash--456dsda45ds",
		PhotoPath: "/userPhoto/1",
		Role: 0,
		IsActive: true,

		Articles: articleList {

		},
		BoughtProjects: []int{3},
	},
	{
		ID: 2,
		FullName: "Necenzūrinis vardas",
		Username: "Nec_name",
		Email: "uzblokuotas@gmail.com",
		PasswordHash: "65d64as8--blockedPassHash--456dsda45ds",
		PhotoPath: "/userPhoto/2",
		Role: 0,
		IsActive: false,

		Articles: articleList {

		},
	},
	{
		ID: 3,
		FullName: "Administratorius",
		Username: "admin",
		Email: "admin@gmail.com",
		PasswordHash: "65d64as8--admin1PassHash--456dsda45ds",
		PhotoPath: "/userPhoto/3",
		Role: 1,
		IsActive: true,

		Articles: articleList {
			article {
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
			article {
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
		},
	},
	{
		ID: 4,
		FullName: "Administratorius2",
		Username: "admin2",
		Email: "admin2@gmail.com",
		PasswordHash: "65d64as8--admin2PassHash--456dsda45ds",
		PhotoPath: "/userPhoto/4",
		Role: 1,
		IsActive: true,

		Articles: articleList {
			article {
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
		},
	},
	{
		ID: 5,
		FullName: "Jonas Jonaitis",
		Username: "JJ",
		Email: "jonas.jonaitis@gmail.com",
		PasswordHash: "65d64as8--user5PassHash--456dsda45ds",
		PhotoPath: "/userPhoto/5",
		Role: 0,
		IsActive: true,

		Articles: articleList {

		},
	},
}

var userIndexer = 5

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

	newUser.ID = userIndexer
	newUser.Role = 0
	newUser.IsActive = true
	newUser.Articles = []article{}
	userIndexer++
	users = append(users, newUser)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
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
		for _, u := range users {
			if u.ID == id {
				return u, http.StatusOK
			}
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
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				return u, http.StatusNoContent
			}
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

		for i, u := range users {
			if u.ID == id {
				u.FullName = updateUser.FullName
				u.Username = updateUser.Username
				u.Email = updateUser.Email
				u.PasswordHash = updateUser.PasswordHash
				u.PhotoPath = updateUser.PhotoPath
				u.Role = updateUser.Role
				u.IsActive = updateUser.IsActive
				remUsers := users[i+1:]
				users = append(users[:i], u)
				users = append(users, remUsers...)
				return u, http.StatusOK
			}
		}

		return user{}, http.StatusNotFound
	}

	return user{}, http.StatusForbidden
}
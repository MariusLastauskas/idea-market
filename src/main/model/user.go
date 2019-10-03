package model

import (
	"encoding/json"
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
}

var userIndexer = 5

func HandleUsersGet(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, role := AuthoriseByToken(r)

	if !isAuthenticated || role != 1 {
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
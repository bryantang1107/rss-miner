package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bryantang1107/Rss_Miner/commons"
	"github.com/bryantang1107/Rss_Miner/internal/database"
	"github.com/bryantang1107/Rss_Miner/models"
	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json: "name"`
	}

	decoder := json.NewDecoder(r.Body) // decode from json to struct
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Failed to Create User", err))
		return
	}

	commons.ResponseWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apiConfig *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	commons.ResponseWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiConfig *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Couldn't retrieve posts: %v", err))
		return
	}
	commons.ResponseWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}

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

func (apiConfig *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json: "name"`
		URL  string `json: "url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Failed to Create User", err))
		return
	}

	commons.ResponseWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiConfig *ApiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
	}

	commons.ResponseWithJSON(w, 201, models.DatabaseFeedsToFeeds(feeds))
}

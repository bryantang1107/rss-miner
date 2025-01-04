package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bryantang1107/Rss_Miner/commons"
	"github.com/bryantang1107/Rss_Miner/internal/database"
	"github.com/bryantang1107/Rss_Miner/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Failed to Create Feed Follow", err))
		return
	}

	commons.ResponseWithJSON(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *ApiConfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiConfig.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Failed to Get Feed Follow", err))
		return
	}
	commons.ResponseWithJSON(w, 201, models.DatabaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiConfig *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id : %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		commons.ResponseWithError(w, 400, fmt.Sprintf("Failed to Get Feed Follow", err))
		return
	}
	commons.ResponseWithJSON(w, 200, struct{}{}) // return empty json obj
}

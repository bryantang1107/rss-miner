package http_handler

import (
	"fmt"
	"net/http"

	"github.com/bryantang1107/Rss_Miner/commons"
	"github.com/bryantang1107/Rss_Miner/internal/auth"
	"github.com/bryantang1107/Rss_Miner/internal/database"
)

// custom type for function signature
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			commons.ResponseWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			commons.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

package http_handler

import (
	"net/http"

	"github.com/bryantang1107/Rss_Miner/commons"
	"github.com/bryantang1107/Rss_Miner/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	commons.ResponseWithJSON(w, 200, struct{}{})
}

func HandleErr(w http.ResponseWriter, r *http.Request) {
	commons.ResponseWithError(w, 400, "Something Went Wrong")
}

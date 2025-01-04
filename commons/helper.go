package commons

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respond with 5XX error:", msg)
	}

	ResponseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed To Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		log.Println("Failed to Write Response")
		return
	}
}

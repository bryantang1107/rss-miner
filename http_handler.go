package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something Went Wrong")
}

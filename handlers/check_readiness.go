package handlers

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "Ready")
}

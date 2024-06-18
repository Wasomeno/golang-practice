package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, error := json.Marshal(payload)
	if error != nil {
		log.Println("Failed to marshal JSON payload:", error)
		w.WriteHeader(500)
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, errorMessage string) {
	respondWithJSON(w, code, map[string]string{"code": fmt.Sprint(code), "error": errorMessage})
}

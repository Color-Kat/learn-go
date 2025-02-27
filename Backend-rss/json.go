package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, error := json.Marshal(payload)

	if error != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println("Responding with 5XX error: ", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

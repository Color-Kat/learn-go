package main

import "net/http"

func handlerReadiness(response http.ResponseWriter, request *http.Request) {
	respondWithJSON(response, 200, struct{}{})
}

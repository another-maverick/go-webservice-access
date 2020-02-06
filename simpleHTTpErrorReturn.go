package main

import "net/http"

func respondWithError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "I am just returning a 403 error", http.StatusForbidden)
}

func main() {
	http.HandleFunc("/", respondWithError)
	http.ListenAndServe(":12345", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyError struct {
	HTTPCode int `json: "-"`
	MyCode int `json: "code, omitempty"`
	Message string `json: "message"`
}

//this function is similar to HTTP error. Difference is that we are sending response in JSON with custom fields
func MyJSONResponse(w http.ResponseWriter, e MyError) {
	// Create an instance of anonymous struct that has error data
	errData := struct{WrapperError MyError `json: "error"`}{e}

	// Convert Error data to JSON
	data, err := json.Marshal(errData)

	// Send generic error if there was a problem in marshalling
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Set HTTP error code in the header
	w.WriteHeader(e.HTTPCode)
	// generate the response. This has application specific error code
	fmt.Fprint(w, string(data))

}

func generateError(w http.ResponseWriter, r *http.Request){
	e := MyError{HTTPCode: http.StatusForbidden,
					MyCode: 12345,
					Message: "My Custom message for code - 12345",}
	MyJSONResponse(w, e)

}

func main() {
	http.HandleFunc("/", generateError)
	http.ListenAndServe(":12345", nil)
}



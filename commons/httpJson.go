package commons

import (
	"encoding/json"
	"net/http"
)

// ReadJSON reads data from incoming HTTP request and marshall it to JSON.
func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// WriteJSON encode any data to a HTTP response writer for sending it as HTTP response.
func WriteJSON(w http.ResponseWriter, statuscode int, data any) {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// WriteError sends a JSON response with a key error.
func WriteError(w http.ResponseWriter, statuscode int, msg string) {
	WriteJSON(w, statuscode, map[string]string{"error": msg})
}

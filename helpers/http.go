package helpers

import "net/http"

// SetHeaders sets header access controls
func SetHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

// EnableJSON sets json as content type in header
func EnableJSON(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

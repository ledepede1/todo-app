package middleware

import "net/http"

func EnableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Handle OPTIONS request
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

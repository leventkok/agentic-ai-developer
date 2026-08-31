package httpapi

import (
	"encoding/json"
	"net/http"
)

type ProblemDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeProblem(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ProblemDetail{Code: status, Message: message})
}
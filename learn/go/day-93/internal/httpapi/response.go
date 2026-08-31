package httpapi

import (
	"encoding/json"
	"net/http"

	"learn/go/day-93/internal/perf/buffer"
)

type ProblemDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	buf := buffer.Get()
	defer buffer.Put(buf)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		writeProblem(w, http.StatusInternalServerError, "failed to encode response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

func writeProblem(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ProblemDetail{Code: status, Message: message})
}

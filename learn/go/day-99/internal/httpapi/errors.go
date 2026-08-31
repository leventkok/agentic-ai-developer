package httpapi

import (
	"context"
	"errors"
	"net/http"

	"learn/go/day-99/internal/domain"
)

func writeDomainError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		writeProblem(w, http.StatusRequestTimeout, err.Error())
		return true
	}
	if errors.Is(err, domain.ErrNotFound) {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
		return true
	}
	if errors.Is(err, domain.ErrForbidden) {
		writeProblem(w, http.StatusForbidden, "forbidden")
		return true
	}
	if domain.IsValidation(err) {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return true
	}
	writeProblem(w, http.StatusInternalServerError, "internal error")
	return true
}

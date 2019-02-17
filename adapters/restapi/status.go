package restapi

import (
	"net/http"
)

var (
	// StatusOK is the HTTP status code of "OK".
	StatusOK = http.StatusOK
	// StatusInternalServerError is the HTTP status code of "Internal Server Error".
	StatusInternalServerError = http.StatusInternalServerError
	// StatusBadRequest is the HTTP status code of "Bad Request"
	StatusBadRequest = http.StatusBadRequest
	// StatusUnprocessableEntity is the HTTP status code of "Unprocessable Entity".
	StatusUnprocessableEntity = http.StatusUnprocessableEntity
)

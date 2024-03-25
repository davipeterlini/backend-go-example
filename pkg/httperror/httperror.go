package httperror

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func SendErrorResponse(w http.ResponseWriter, errResponse *ErrorResponse) {
	w.WriteHeader(errResponse.Code)
	json.NewEncoder(w).Encode(errResponse)
}

func SendErrorSpecificResponse(w http.ResponseWriter, code int, message string) {
	response := NewErrorResponse(code, message)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func HandleNotFoundError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusNotFound, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleBadRequestError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusBadRequest, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleInternalServerError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusInternalServerError, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleConflictError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusConflict, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleUnauthorizedError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusUnauthorized, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleForbiddenError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusForbidden, errMsg)
	SendErrorResponse(w, errResponse)
}

func HandleMethodNotAllowedError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusMethodNotAllowed, errMsg)
	SendErrorResponse(w, errResponse)
}

// TODO - Handlers For CRUD - fix messages
func HandleCreateError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusBadRequest, "Create operation failed: "+err.Error())
}

func HandleUpdateError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusUnprocessableEntity, "Update operation failed: "+err.Error())
}

func HandleDeleteError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusNotFound, "Delete operation failed: "+err.Error())
}

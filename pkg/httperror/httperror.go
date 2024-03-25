// pkg/httperror/httperror.go
package httperror

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse.
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// SendErrorResponse sends an error response
func SendErrorSpecificResponse(w http.ResponseWriter, code int, message string) {
	response := NewErrorResponse(code, message)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

// SendErrorResponse sends an error response.
func SendErrorResponse(w http.ResponseWriter, errResponse *ErrorResponse) {
	w.WriteHeader(errResponse.Code)
	json.NewEncoder(w).Encode(errResponse)
}

// HandleNotFoundError handles not-found errors.
func HandleNotFoundError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusNotFound, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleBadRequestError handles bad-request errors.
func HandleBadRequestError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusBadRequest, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleInternalServerError handles internal-server errors.
func HandleInternalServerError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusInternalServerError, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleConflictError handles conflict errors.
func HandleConflictError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusConflict, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleUnauthorizedError handles unauthorized errors.
func HandleUnauthorizedError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusUnauthorized, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleForbiddenError handles forbidden errors.
func HandleForbiddenError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusForbidden, errMsg)
	SendErrorResponse(w, errResponse)
}

// HandleMethodNotAllowedError handles HTTP errors for unsupported methods.
func HandleMethodNotAllowedError(w http.ResponseWriter, errMsg string) {
	errResponse := NewErrorResponse(http.StatusMethodNotAllowed, errMsg)
	SendErrorResponse(w, errResponse)
}

// Handlers For CRUD

func HandleCreateError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusBadRequest, "Create operation failed: "+err.Error())
}

// HandleUpdateError handles errors during update operations
func HandleUpdateError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusUnprocessableEntity, "Update operation failed: "+err.Error())
}

func HandleDeleteError(w http.ResponseWriter, err error) {
	SendErrorSpecificResponse(w, http.StatusNotFound, "Delete operation failed: "+err.Error())
}

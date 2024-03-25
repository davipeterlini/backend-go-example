package httperror

import (
	"net/http"
)

// ErrorHandlerInterceptor intercepts responses for common HTTP error conditions.
func ErrorHandlerInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use a ResponseWriter that allows us to intercept and inspect the status code
		interceptWriter := NewInterceptWriter(w)
		next.ServeHTTP(interceptWriter, r)

		// After the handler has finished, check the status code and handle errors
		switch interceptWriter.statusCode {
		case http.StatusNotFound:
			HandleNotFoundError(w, "The requested resource was not found")
		case http.StatusBadRequest:
			HandleBadRequestError(w, "Invalid request payload")
		case http.StatusInternalServerError:
			HandleInternalServerError(w, "Internal Server Error")
		case http.StatusConflict:
			HandleConflictError(w, "Conflict")
		case http.StatusForbidden:
			HandleForbiddenError(w, "Forbidden")
		case http.StatusMethodNotAllowed:
			HandleMethodNotAllowedError(w, "The request method is not supported for the requested resource")
		}
	})
}

// NewInterceptWriter creates an instance of InterceptWriter.
func NewInterceptWriter(w http.ResponseWriter) *InterceptWriter {
	return &InterceptWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, // Default to 200 OK
	}
}

// InterceptWriter is a custom http.ResponseWriter that captures the status code.
type InterceptWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the underlying WriteHeader.
func (iw *InterceptWriter) WriteHeader(code int) {
	iw.statusCode = code
	iw.ResponseWriter.WriteHeader(code)
}

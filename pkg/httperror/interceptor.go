package httperror

import (
	"net/http"
)

func ErrorHandlerInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interceptWriter := NewInterceptWriter(w)
		next.ServeHTTP(interceptWriter, r)

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

func NewInterceptWriter(w http.ResponseWriter) *InterceptWriter {
	return &InterceptWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

type InterceptWriter struct {
	http.ResponseWriter
	statusCode int
}

func (iw *InterceptWriter) WriteHeader(code int) {
	iw.statusCode = code
	iw.ResponseWriter.WriteHeader(code)
}

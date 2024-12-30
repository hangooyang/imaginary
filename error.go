package imaginary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrNotFound             = NewError("Not found", http.StatusNotFound)
	ErrInvalidAPIKey        = NewError("Invalid or missing API key", http.StatusUnauthorized)
	ErrMethodNotAllowed     = NewError("HTTP method not allowed. Try with a POST or GET method (-enable-url-source flag must be defined)", http.StatusMethodNotAllowed)
	ErrGetMethodNotAllowed  = NewError("GET method not allowed. Make sure remote URL source is enabled by using the flag: -enable-url-source", http.StatusMethodNotAllowed)
	ErrUnsupportedMedia     = NewError("Unsupported media type", http.StatusNotAcceptable)
	ErrOutputFormat         = NewError("Unsupported output image format", http.StatusBadRequest)
	ErrEmptyBody            = NewError("Empty or unreadable image", http.StatusBadRequest)
	ErrMissingParamFile     = NewError("Missing required param: file", http.StatusBadRequest)
	ErrInvalidFilePath      = NewError("Invalid file path", http.StatusBadRequest)
	ErrInvalidImageURL      = NewError("Invalid image URL", http.StatusBadRequest)
	ErrMissingImageSource   = NewError("Cannot process the image due to missing or invalid params", http.StatusBadRequest)
	ErrNotImplemented       = NewError("Not implemented endpoint", http.StatusNotImplemented)
	ErrInvalidURLSignature  = NewError("Invalid URL signature", http.StatusBadRequest)
	ErrURLSignatureMismatch = NewError("URL signature mismatch", http.StatusForbidden)
	ErrResolutionTooBig     = NewError("Image resolution is too big", http.StatusUnprocessableEntity)
)

type Error struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"status"`
}

func (e Error) JSON() []byte {
	buf, _ := json.Marshal(e)
	return buf
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) HTTPCode() int {
	if e.Code >= 400 && e.Code <= 511 {
		return e.Code
	}
	return http.StatusServiceUnavailable
}

func NewError(err string, code int) Error {
	err = strings.Replace(err, "\n", "", -1)
	return Error{Message: err, Code: code}
}

func sendErrorResponse(w http.ResponseWriter, httpStatusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	_, _ = w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\", \"status\": %d}", err.Error(), httpStatusCode)))
}

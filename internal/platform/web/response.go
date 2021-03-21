package web

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, statusCode int) error {
	d, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.Wrap(err, "error marshalling product")
	}

	w.Header().Set("content-type", "application/json; charset=urf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(d); err != nil {
		return errors.Wrap(err, "error writing")
	}

	return nil
}

func RespondError(w http.ResponseWriter, err error) error {

	if webErr, ok := err.(*Error); ok {
		resp := ErrorResponse{
			Error: webErr.Err.Error(),
		}

		return Respond(w, resp, webErr.Status)

	}

	resp := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}

	return Respond(w, resp, http.StatusInternalServerError)

}

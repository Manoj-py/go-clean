package httpUtils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"reflect"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJsonAndValidate(w http.ResponseWriter, r *http.Request, data any) error {
	if err := readJsonFromBody(w, r, data); err != nil {
		return err
	}

	if err := validation.ValidateRequest(data); err != nil {
		return err
	}

	return nil
}

func readJsonFromBody(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 10 << 20
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		log.Println(err)
		log.Println(reflect.TypeOf(err))
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("there is no in the JSON value in the request")
	}

	return nil
}

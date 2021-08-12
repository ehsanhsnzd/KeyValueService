package Services

import (
	"biges/Models"
	"biges/Repository"
	"errors"
	"net/http"
)

type SERVICE struct {
	DataRepository *Repository.Data
}

// create object
func (service *SERVICE) create(r *http.Request, w http.ResponseWriter,content Models.Content) error {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if validate(key, value) {
		result := content.SetValue(key, value)
		if result != false {
			return ReturnSuccessResponse(w)
		}
	}

	return errors.New("cannot assign new key")
}

// get object
func (service *SERVICE) get(r *http.Request, w http.ResponseWriter,content Models.Content) error {
	key := r.URL.Query().Get("key")

	if validate(key) {
		value, result := content.GetValue(key)
		if result != false {
			return ReturnSuccessDataResponse(w, value)
		}
	}

	return errors.New("key not found")
}

// validate for uri queries
func validate(values... string) bool {
	for _, value := range values {
		if value == "" {
			return false
		}
	}
	return true
}

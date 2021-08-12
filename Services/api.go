package Services

import (
	"fmt"
	"net/http"
)

func (service *SERVICE)HandleSet(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.URL.Path != "/set" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		err = service.create(r, w, service.DataRepository)
	default:
		fmt.Fprintf(w, "Sorry, only GET method are supported.")
	}

	if err != nil {
		ReturnErrorResponse(w)
		return
	}

}


func (service *SERVICE)HandleGet(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.URL.Path != "/get" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = service.get(r, w, service.DataRepository)
	default:
		fmt.Fprintf(w, "Sorry, only POST method are supported.")
	}

	if err != nil {
		ReturnErrorResponse(w)
		return
	}

}

func ReturnErrorResponse(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"error\"}"))
}

func ReturnSuccessResponse(w http.ResponseWriter) error {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"success\"}"))
	return nil
}

func ReturnSuccessDataResponse(w http.ResponseWriter, data string) error {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"success\",\"data\": " + data + "}"))
	return nil
}

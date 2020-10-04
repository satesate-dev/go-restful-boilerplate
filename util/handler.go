package util

import (
	"encoding/json"
	"net/http"
)

type (
	HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)
	Response    struct {
		BaseResponse
		Data interface{} `json:"data"`
	}
	BaseResponse struct {
		Errors []string `json:"errors,omitempty"`
	}
)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errs []string

	r.ParseForm()

	data, err := fn(w, r)
	if err != nil {
		errs = append(errs, err.Error())
		//w.WriteHeader(err)
	}

	resp := Response{
		Data: data,
		BaseResponse: BaseResponse{
			Errors: errs,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return
	}
}

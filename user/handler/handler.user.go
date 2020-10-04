package handler

import (
	"net/http"

	"github.com/satesate-dev/go-restful-boilerplate/user/api"
)

func NewHandlerUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	return api.GetAllUser(ctx)
}

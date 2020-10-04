package endpoint

import (
	"net/http"

	"github.com/satesate-dev/go-restful-boilerplate/util"

	"github.com/satesate-dev/go-restful-boilerplate/user/handler"

	"github.com/gorilla/mux"
)

func NewUserEndpoint(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	users.Handle("", util.HandlerFunc(handler.NewHandlerUser)).Methods(http.MethodGet)
}

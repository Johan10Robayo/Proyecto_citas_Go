package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)

	http.Handle("/", router)

	return router
}

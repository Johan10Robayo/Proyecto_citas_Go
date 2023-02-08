package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"healthpriority.com/handlers"
)

func Init() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/registrar", handlers.RegistroCliente).Methods("POST")
	router.Use(handlers.JwtMiddleware)
	router.HandleFunc("/api/agendageneral", handlers.AgendarMedicoG).Methods("POST")
	http.Handle("/", router)

	return router
}

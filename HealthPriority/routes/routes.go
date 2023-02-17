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
	router.HandleFunc("/api/autorizacion", handlers.Autorizacion).Methods("POST")
	router.HandleFunc("/api/getAutorizaciones", handlers.GetAutorizaciones).Methods("POST")
	http.Handle("/", router)

	return router
}

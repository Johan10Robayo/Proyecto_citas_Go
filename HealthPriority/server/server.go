package server

import (
	"log"
	"net/http"

	"healthpriority.com/routes"
)

func Init() {
	servidor := http.Server{
		Addr:    ":8080",
		Handler: routes.Mux,
	}
	log.Println("Listening.. at http://localhost:8080/index")
	err := servidor.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

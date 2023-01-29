package routes

import (
	"net/http"

	"healthpriority.com/handlers"
)

var Mux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("/index", handlers.Index)

}

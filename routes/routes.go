package routes

import (
	"api/rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux" // <--- 1. A importação que faltava
)

func HandleRequests() {
	r := mux.NewRouter()

	// 2. Correção abaixo: removemos o ".http" extra
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

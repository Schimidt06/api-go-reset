package routes

import (
	"api/rest/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

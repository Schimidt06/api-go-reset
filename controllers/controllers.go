package controllers

import (
	"api/rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter , r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for _ , personalidade := range models.Personalidades {
        if strconv.Itoa(personalidade.Id) == id {
            json.NewEncoder(w).Encode(personalidade)
        }
    }
}



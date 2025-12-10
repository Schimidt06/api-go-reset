package main

import (
	"api/rest/models"
	"api/rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Albert Einstein", Descricao: "Físico teórico conhecido pela teoria da relatividade."},
		{Nome: "Marie Curie", Descricao: "Pioneira na pesquisa sobre radioatividade."},
		{Nome: "Isaac Newton", Descricao: "Matemático e físico que formulou as leis do movimento."},
	}

	fmt.Println("Hello, World!")
	routes.HandleRequests()
}

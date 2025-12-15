package main

import (
	"api/rest/database"
	"api/rest/models"
	"api/rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico conhecido pela teoria da relatividade."},
		{Id: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade."},
		{Id: 3, Nome: "Isaac Newton", Historia: "Matemático e físico que formulou as leis do movimento."},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Hello, World!")
	routes.HandleRequests()
}

package routes

import (
	"api/rest/controllers"
	"api/rest/middleware" // Importa o nosso novo pacote de middleware
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquests() {
	// Inicializa o roteador do Gorilla Mux
	r := mux.NewRouter()

	// [NOVO] Registra o Middleware
	// O comando r.Use diz ao roteador: "Antes de executar qualquer função abaixo,
	// passe por essa função primeiro".
	// Isso garante que TODAS as respostas tenham o "Content-Type: application/json".
	r.Use(middleware.ContentTypeMiddleware)

	// Rota Home (geralmente usada para verificar se a API está online)
	r.HandleFunc("/", controllers.Home)

	// Rotas da API
	// .Methods("Verb") restringe a rota para aceitar apenas aquele método HTTP específico.
	// Se alguém tentar fazer um POST na rota de GET, o Mux bloqueia automaticamente.

	// Busca todas
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")

	// Busca uma específica (usa {id} como variável)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")

	// Cria uma nova
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")

	// Deleta
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")

	// Edita
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods("Put")

	// Inicia o servidor na porta 8000
	// log.Fatal envolve o comando para que, se o servidor cair ou falhar ao iniciar,
	// o erro seja mostrado no terminal imediatamente.
	log.Fatal(http.ListenAndServe(":8000", r))
}
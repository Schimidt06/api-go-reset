package routes

import (
    "api/rest/controllers" // Importa a lógica (as funções que criamos antes)
    "log"                  // Lib para logar erros no terminal
    "net/http"             // Lib para subir o servidor web

    "github.com/gorilla/mux" // O Roteador (Router)
)

// HandleRequests: Função pública que inicia o servidor e define as rotas.
// OBS: Notei um pequeno erro de digitação no nome: 'HandleResquests' (tem um 's' extra).
// O ideal seria 'HandleRequests'.
func HandleResquests() {
    // 1. Instanciando o Roteador:
    // O 'mux' é mais poderoso que o roteador padrão do Go.
    // Ele permite capturar variáveis na URL (ex: {id}) e restringir verbos (GET, POST).
    r := mux.NewRouter()

    // 2. Rota Home (Página Inicial):
    // Quando acessar "localhost:8000/", chama a função Home.
    r.HandleFunc("/", controllers.Home)

    // 3. Rota de Listagem (GET ALL):
    // .Methods("Get"): Isso é CRUCIAL. Garante que essa rota só aceite requisições de leitura.
    // Se alguém tentar enviar dados (POST) para cá, receberá erro "405 Method Not Allowed".
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")

    // 4. Rota de Busca Única (GET BY ID):
    // A sintaxe "{id}" diz ao Mux: "O que vier aqui é uma variável, não texto fixo".
    // O controller vai ler isso usando mux.Vars(r)["id"].
    r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")

    // 5. Rota de Criação (POST):
    // Mapeamos a mesma URL da listagem, mas mudamos o VERBO para "Post".
    // Isso é o padrão RESTful: Mesma URL, intenções diferentes baseadas no verbo HTTP.
    r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")

    // 6. Rota de Deleção (DELETE):
    r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")

    // 7. Subindo o Servidor (O Loop Infinito):
    // http.ListenAndServe(":8000", r) abre a porta 8000 e fica escutando para sempre.
    // O log.Fatal envolve tudo: se o servidor cair (ex: porta 8000 já está em uso),
    // ele mata o programa e mostra o erro no terminal.
    log.Fatal(http.ListenAndServe(":8000", r))
}
package routes

import (
    "api/rest/controllers"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func HandleResquests() { // Lembra de arrumar o nome para HandleRequests depois ;)
    r := mux.NewRouter()

    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
    r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
    r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
    r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")

    // [CORREÇÃO CRÍTICA] Movi esta linha para CIMA.
    // 8. Rota de Edição (PUT):
    // O verbo PUT é o padrão mundial para "Atualizar um recurso existente".
    // Diferença do POST: POST cria novo. PUT substitui o que já existe.
    // Ele precisa do {id} na URL para saber QUEM ele vai atualizar.
    r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods("Put")

    // 9. Subindo o Servidor (O Ponto Sem Volta):
    // ATENÇÃO SCHIMIDT: Esta linha TEM que ser a última.
    // O programa "trava" aqui e fica rodando. Nada escrito abaixo dessa linha será executado.
    log.Fatal(http.ListenAndServe(":8000", r))
}
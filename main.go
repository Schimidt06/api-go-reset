package main

import (
    "api/rest/database" // Importa o inicializador do Banco
    "api/rest/models"   // Importa a estrutura de dados
    "api/rest/routes"   // Importa o roteador
    "fmt"
)

func main() {
    // 1. [ALERTA DE CÓDIGO MORTO / LEGADO]
    // Schimidt, lembra que configuramos o Docker e o Postgres?
    // Aqui você está criando dados na Memória RAM (Slice).
    // O problema: Seus Controllers (no arquivo controllers.go) estão lendo do BANCO (database.DB.Find).
    // Resultado: A API vai ignorar completamente o Einstein, Marie Curie e Newton que você escreveu aqui.
    // Ela vai mostrar o "Deodato" e a "Carmela" que estão no seu script SQL do Docker.
    //
    // Solução Profissional: Podemos apagar esse bloco inteiro. O banco é quem manda agora.
    models.Personalidades = []models.Personalidade{
        {Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico conhecido pela teoria da relatividade."},
        {Id: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade."},
        {Id: 3, Nome: "Isaac Newton", Historia: "Matemático e físico que formulou as leis do movimento."},
    }

    // 2. O Handshake (Aperto de Mão):
    // Essa linha é vital. Ela vai até o database/db.go e tenta abrir a porta 5432.
    // Se o Docker não estiver rodando, o programa vai dar "Panic" e fechar aqui mesmo.
    database.ConectaComBancoDeDados()

    // 3. Log de Inicialização:
    // É sempre bom ter um feedback visual no terminal para saber que o programa não travou.
    fmt.Println("Iniciando o servidor Rest com Go...")

    // 4. O Loop Infinito:
    // Chama o arquivo routes.go.
    // ATENÇÃO: Lembre-se que lá estava escrito 'HandleResquests' (com erro de digitação).
    // O Go vai travar a execução aqui e ficar escutando a porta 8000 eternamente.
    routes.HandleResquests()
}
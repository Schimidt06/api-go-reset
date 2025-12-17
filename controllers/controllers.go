package controllers

import (
    "api/rest/database" // Importa nossa conexão com o banco (GORM)
    "api/rest/models"   // Importa a estrutura (Struct) da tabela Personalidade
    "encoding/json"     // Lib nativa para traduzir Go <-> JSON
    "fmt"
    "net/http"          // Lib nativa para gerenciar requisições e respostas web

    "github.com/gorilla/mux" // Roteador poderoso para capturar variáveis na URL
)

// Home: A porta de entrada da API.
// w (ResponseWriter): É onde "escrevemos" a resposta que vai para o navegador/cliente.
// r (Request): Contém todos os dados que chegaram (cabeçalhos, corpo, IP, etc).
func Home(w http.ResponseWriter, r *http.Request) {
    // Fprintf escreve uma string formatada diretamente na saída HTTP.
    fmt.Fprintln(w, "Welcome to the Home Page!")
}

// TodasPersonalidades: Busca TUDO no banco (equivalente a um SELECT *).
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    // 1. Criação do Slice (Lista):
    // Em Go, 'var p []...' cria um slice vazio que vai receber vários registros.
    // É como um TList ou Array dinâmico no Delphi/C#.
    var p []models.Personalidade

    // 2. A Mágica do Ponteiro (&):
    // Passamos '&p' (o endereço de memória) para o GORM.
    // O GORM precisa do endereço para preencher a variável original com os dados do banco.
    database.DB.Find(&p)

    // 3. Serialização (Marshalling):
    // Cria um Encoder que joga dados na resposta (w) e codifica o slice 'p' para JSON.
    json.NewEncoder(w).Encode(p)
}

// RetornaUmaPersonalidade: Busca um único registro baseado no ID da URL.
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    // 1. Captura de Variáveis:
    // O Mux extrai os parâmetros definidos na rota (ex: /api/personalidades/{id}).
    vars := mux.Vars(r)
    id := vars["id"]

    // 2. Instância Única:
    // Aqui criamos uma struct vazia, não um slice, pois esperamos apenas 1 resultado.
    var Personalidade models.Personalidade

    // 3. Busca com Filtro (WHERE id = ?):
    // O método .First() busca o primeiro registro que coincida com o 'id' passado.
    // Novamente, passamos o endereço de memória (&Personalidade) para ser preenchido.
    database.DB.First(&Personalidade, id)

    // 4. Retorno:
    json.NewEncoder(w).Encode(Personalidade)
}

// CriaUmaNovaPersonalidade: Recebe JSON do cliente e salva no banco.
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
    var novaPersonalidade models.Personalidade

    // 1. Deserialização (Unmarshalling):
    // NewDecoder pega o 'r.Body' (o JSON cru que veio na requisição).
    // .Decode(&...) converte esse JSON para a struct Go e preenche a variável.
    json.NewDecoder(r.Body).Decode(&novaPersonalidade)

    // 2. Persistência:
    // O comando .Create() gera o INSERT no banco de dados.
    database.DB.Create(&novaPersonalidade)

    // 3. Confirmação:
    // Devolvemos o objeto criado (agora com ID gerado) para confirmar que deu certo.
    json.NewEncoder(w).Encode(novaPersonalidade)
}

// DeletaUmaPersonalidade: Remove um registro baseado no ID.
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    // Captura o ID da URL
    vars := mux.Vars(r)
    id := vars["id"]

    var Personalidade models.Personalidade

    // O comando .Delete() executa o DELETE FROM personalidades WHERE id = X.
    // O GORM precisa da struct para saber qual tabela usar, e do ID para saber quem apagar.
    database.DB.Delete(&Personalidade, id)

    // Retorna o objeto (geralmente vazio ou com dados residuais) como confirmação técnica.
    json.NewEncoder(w).Encode(Personalidade)
}
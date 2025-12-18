package controllers

import (
	"api/rest/database" // Camada de acesso ao banco de dados (GORM já configurado)
	"api/rest/models"   // Modelos que representam as tabelas do banco (Structs)
	"encoding/json"     // Biblioteca padrão para conversão entre Go <-> JSON
	"fmt"
	"net/http"          // Base do servidor HTTP: Request, Response, Status Codes

	"github.com/gorilla/mux" // Roteador HTTP para rotas dinâmicas (path params)
)


// Home
// Responsável por responder a rota raiz da API.
// Normalmente usada como health check ou endpoint inicial.
func Home(w http.ResponseWriter, r *http.Request) {
	// Escreve diretamente na resposta HTTP
	// Ideal para testes rápidos de funcionamento da API
	fmt.Fprintln(w, "Welcome to the Home Page!")
}


// TodasPersonalidades
// Retorna todos os registros da tabela Personalidade.
// Equivalente ao comando SQL: SELECT * FROM personalidades
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	// Slice que armazenará múltiplos registros retornados do banco
	// Slice é dinâmico e cresce conforme o GORM popula os dados
	var personalidades []models.Personalidade

	// Executa a consulta no banco
	// Passamos o endereço de memória (&) para que o GORM consiga
	// preencher a variável original com os dados retornados
	database.DB.Find(&personalidades)

	// Converte automaticamente o slice Go em JSON
	// e escreve o resultado na resposta HTTP
	json.NewEncoder(w).Encode(personalidades)
}


// RetornaUmaPersonalidade
// Busca um único registro com base no ID informado na URL.
// Exemplo de rota: /api/personalidades/{id}
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// Captura os parâmetros dinâmicos da rota
	// O mux transforma {id} da URL em um map
	vars := mux.Vars(r)
	id := vars["id"]

	// Struct única, pois esperamos apenas um registro
	var personalidade models.Personalidade

	// Consulta no banco buscando o primeiro registro
	// cujo ID corresponda ao valor informado
	// Internamente gera: SELECT * FROM personalidades WHERE id = ? LIMIT 1
	database.DB.First(&personalidade, id)

	// Retorna o objeto encontrado em formato JSON
	json.NewEncoder(w).Encode(personalidade)
}


// CriaUmaNovaPersonalidade
// Recebe um JSON via corpo da requisição e cria um novo registro no banco.
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// Struct que irá receber os dados enviados pelo cliente
	var novaPersonalidade models.Personalidade

	// Decodifica o JSON do corpo da requisição
	// e preenche automaticamente a struct Go
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	// Persiste os dados no banco
	// Gera internamente um INSERT
	database.DB.Create(&novaPersonalidade)

	// Retorna o objeto salvo, geralmente já contendo ID gerado pelo banco
	json.NewEncoder(w).Encode(novaPersonalidade)
}


// DeletaUmaPersonalidade
// Remove um registro do banco com base no ID informado na URL.
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// Captura o ID vindo da rota
	vars := mux.Vars(r)
	id := vars["id"]

	// Struct necessária para que o GORM saiba qual tabela utilizar
	var personalidade models.Personalidade

	// Executa o DELETE
	// Internamente: DELETE FROM personalidades WHERE id = ?
	database.DB.Delete(&personalidade, id)

	// Retorno técnico para confirmar a operação
	json.NewEncoder(w).Encode(personalidade)
}


// EditaUmaPersonalidade
// Atualiza um registro existente com base no ID.
func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// Captura o ID da URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Busca o registro existente no banco
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	// Sobrescreve os campos da struct com os dados recebidos no JSON
	json.NewDecoder(r.Body).Decode(&personalidade)

	// Salva as alterações
	// O GORM identifica que é um UPDATE pois o ID já existe
	database.DB.Save(&personalidade)

	// Retorna o objeto atualizado
	json.NewEncoder(w).Encode(personalidade)
}
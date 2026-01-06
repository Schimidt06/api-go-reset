package controllers

import (
	"api/rest/database" // Importa a configuração do banco de dados
	"api/rest/models"   // Importa os modelos (structs) que representam as tabelas
	"encoding/json"     // Ferramenta nativa do Go para lidar com JSON (Encoder/Decoder)
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // Biblioteca externa para manipular rotas e variáveis de URL
)

// Home exibe uma mensagem simples.
// Útil para verificar se a API está no ar (Health Check).
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// TodasPersonalidades busca todos os registros no banco.
// Verbo HTTP esperado: GET
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	// Cria um slice (lista) de Personalidade para receber os dados
	var p []models.Personalidade

	// O GORM executa algo como "SELECT * FROM personalidades" e preenche 'p'
	database.DB.Find(&p)

	// Codifica a lista 'p' em JSON e escreve na resposta (w)
	json.NewEncoder(w).Encode(p)
}

// RetornaUmaPersonalidade busca um registro específico pelo ID.
// Verbo HTTP esperado: GET
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) captura as variáveis da URL (ex: id em /api/personalidades/{id})
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	// O GORM busca o primeiro registro onde o ID bate com a variável.
	// SQL equivalente: SELECT * FROM personalidades WHERE id = ? LIMIT 1
	database.DB.First(&personalidade, id)

	// Retorna o objeto encontrado como JSON
	json.NewEncoder(w).Encode(personalidade)
}

// CriaUmaNovaPersonalidade recebe dados JSON e salva no banco.
// Verbo HTTP esperado: POST
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	// O Decoder lê o corpo da requisição (r.Body) e converte o JSON para a Struct do Go
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	// O GORM pega a struct preenchida e salva no banco.
	// SQL equivalente: INSERT INTO personalidades (nome, historia) VALUES (...)
	database.DB.Create(&novaPersonalidade)

	// Retorna o objeto criado (agora com o ID gerado pelo banco) para confirmar
	json.NewEncoder(w).Encode(novaPersonalidade)
}

// DeletaUmaPersonalidade remove um registro pelo ID.
// Verbo HTTP esperado: DELETE
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	// Remove o registro baseando-se no ID.
	// SQL equivalente: DELETE FROM personalidades WHERE id = ?
	database.DB.Delete(&personalidade, id)

	// Retorna o objeto vazio ou dados da operação (depende da implementação do GORM)
	json.NewEncoder(w).Encode(personalidade)
}

// EditaUmaPersonalidade atualiza os dados de um registro existente.
// Verbo HTTP esperado: PUT
func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	// 1º Passo: Busca a personalidade no banco para garantir que ela existe e obter o ID dela na struct
	database.DB.First(&personalidade, id)

	// 2º Passo: Lê o JSON enviado pelo usuário e atualiza os campos da struct 'personalidade'
	json.NewDecoder(r.Body).Decode(&personalidade)

	// 3º Passo: Salva as alterações no banco.
	// Como a struct já tem um ID (do passo 1), o GORM sabe que deve fazer um UPDATE e não um INSERT
	database.DB.Save(&personalidade)

	// Retorna a personalidade já editada
	json.NewEncoder(w).Encode(personalidade)
}
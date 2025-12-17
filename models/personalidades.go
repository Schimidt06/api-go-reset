package models

// Personalidade define a estrutura de dados (o 'Shape') do nosso objeto.
// Em C# isso seria um DTO (Data Transfer Object) ou uma Class Model.
// Em Delphi isso seria um Type Record.
type Personalidade struct {
    // 1. O Campo 'Id':
    // O campo começa com letra MAIÚSCULA (Id) para ser PÚBLICO dentro do Go.
    // A tag `json:"id"` avisa: "Quando gerar o JSON para o navegador, chame isso de 'id' minúsculo".
    Id int `json:"id"`

    // 2. O Campo 'Nome':
    // Mesma lógica. No Go é 'Nome', no JSON vira 'nome'.
    Nome string `json:"nome"`

    // 3. [CORREÇÃO CRÍTICA] O Campo 'Historia':
    // Schimidt, notei um erro aqui no seu código original: `json:"Historia "`.
    // Tinha um espaço em branco no final e a letra H maiúscula.
    // Isso faria o JSON retornar chave "Historia " (com espaço), quebrando quem for consumir a API.
    // Padronizei para minúsculo e sem espaço.
    Historia string `json:"historia"`
}

// 4. [DÍVIDA TÉCNICA] Slice Global (Array Dinâmico):
// Esta variável cria uma lista em memória RAM.
// ATENÇÃO: Agora que você configurou o banco de dados (Postgres),
// essa variável provavelmente se tornou inútil (código morto),
// pois os dados virão do banco, e não dessa variável.
var Personalidades []Personalidade
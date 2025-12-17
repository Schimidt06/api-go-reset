package database

import (
    "gorm.io/driver/postgres" // Driver específico para o PostgreSQL (o "tradutor" do banco)
    "gorm.io/gorm"            // O ORM em si (quem facilita o CRUD)
)

// Bloco de declaração de variáveis globais
var (
    // [IMPORTANTE] Visibilidade (Scope):
    // Em Go, a primeira letra MAIÚSCULA define que a variável é PÚBLICA (Exported).
    // Por isso 'DB' é maiúsculo: precisamos acessá-lo lá no arquivo 'controllers'.
    // O asterisco (*) indica que é um PONTEIRO. Não queremos uma cópia do banco,
    // queremos apontar para a MESMA conexão ativa em toda a aplicação.
    DB *gorm.DB

    // 'err' é minúsculo, então é PRIVADA. Só esse arquivo 'database.go' enxerga ela.
    err error
)

func ConectaComBancoDeDados() {
    // 1. String de Conexão (DSN - Data Source Name):
    // Aqui definimos as credenciais.
    // host=localhost: O banco está na sua máquina.
    // sslmode=disable: Desativa criptografia (ok para desenvolvimento local, perigoso em produção).
    stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

    // 2. Abertura da Conexão:
    // gorm.Open tenta bater na porta 5432.
    // Ele retorna DOIS valores: a conexão (jogada em DB) e um erro (jogado em err), se houver.
    DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

    // 3. Tratamento de Erro Fatal:
    // Go não usa "try/catch" clássico. Verificamos se 'err' é diferente de nulo.
    if err != nil {
        // Panic: É o "Kill Switch".
        // Ele para a execução do programa IMEDIATAMENTE e mostra o erro no console.
        // Usamos panic aqui porque, sem banco, a API não serve para nada.
        panic("Não foi possível conectar com o banco de dados: " + err.Error())
    }
}
package middleware

import "net/http"

// ContentTypeMiddleware
// Essa função funciona como uma "camada extra" de proteção e configuração.
// Ela recebe o próximo handler (a rota que seria chamada) e retorna um novo handler com a configuração aplicada.
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// Antes de chamar o controller, definimos o cabeçalho padrão.
		// Isso garante que TODAS as rotas que usarem esse middleware respondam como "application/json".
		w.Header().Set("Content-Type", "application/json")

		// "next" é a próxima função na fila (o seu controller, por exemplo: TodasPersonalidades).
		// O ServeHTTP passa a bola para ela. Sem isso, a requisição trava aqui.
		next.ServeHTTP(w, r)
	})
}
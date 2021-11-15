package rotas

import "net/http"

// Rota representa todas as rotas acessiveis da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

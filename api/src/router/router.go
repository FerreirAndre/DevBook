package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := rotas.Configurar(mux.NewRouter())

	return r
}

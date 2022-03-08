package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// go get github.com/gorilla/mux
// go get github.com/go-sql-driver/mysql
// go get golang.org/x/crypto/bcrypt
// go get github.com/badoux/checkmail

func main() {
	config.Carregar()

	fmt.Println("API Iniciada", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router.Gerar()))
}

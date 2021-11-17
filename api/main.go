package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println("API Iniciada")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router.Gerar()))
}

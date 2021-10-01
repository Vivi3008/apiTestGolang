package main

import (
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases"
	api "github.com/Vivi3008/apiTestGolang/http"
	"github.com/Vivi3008/apiTestGolang/store"
)

const addr = ":3000"

func main() {
	accountStore := store.NewAccountStore()
	accountsUsecase := usecases.CreateNewAccount(accountStore)

	server := api.NewServer(accountsUsecase)

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}

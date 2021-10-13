package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/usecases"
	api "github.com/Vivi3008/apiTestGolang/http"
	"github.com/Vivi3008/apiTestGolang/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	addr := os.Getenv("PORT")

	accountStore := store.NewAccountStore()
	transStore := store.NewTransferStore()
	transferStore := usecases.SaveNewTransfer(transStore)
	accountsUsecase := usecases.CreateNewAccount(accountStore)

	server := api.NewServer(accountsUsecase, transferStore)

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}

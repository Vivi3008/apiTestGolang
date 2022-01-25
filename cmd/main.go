package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	api "github.com/Vivi3008/apiTestGolang/gateways/http"
	"github.com/Vivi3008/apiTestGolang/store"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	connect, err := postgres.ConnectPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, connect, err)
		os.Exit(1)
	}

	addr := os.Getenv("PORT")

	accountStore := store.NewAccountStore()
	transStore := store.NewTransferStore()
	billStore := store.NewBillStore()

	accUsecase := account.NewAccountUsecase(accountStore)
	transferStore := transfers.NewTransferUsecase(transStore, accUsecase)
	blStore := bill.NewBillUseCase(billStore, accUsecase)

	server := api.NewServer(accUsecase, transferStore, blStore)

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}

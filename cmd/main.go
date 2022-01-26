package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/commom/config"
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
		sendError(err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		sendError(err)
	}

	_, err = postgres.ConnectPool(cfg)
	if err != nil {
		sendError(err)
	}

	addr := cfg.API.Port

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

func sendError(err error) {
	log.Fatalf("Error loading database config: %s", err)
	os.Exit(1)
}

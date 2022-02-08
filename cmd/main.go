package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/commom/config"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"
	account_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	transfers_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"
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

	ctx := context.Background()
	db, err := postgres.ConnectPool(ctx, cfg)
	if err != nil {
		sendError(err)
	}
	defer db.Close()

	accountStore := account_postgres.NewRepository(db)
	transStore := transfers_postgres.NewRepository(db)
	billStore := store.NewBillStore()

	accUsecase := account.NewAccountUsecase(accountStore)
	transferStore := transfers.NewTransferUsecase(transStore, accUsecase)
	blStore := bill.NewBillUseCase(billStore, accUsecase)

	server := api.NewServer(accUsecase, transferStore, blStore)

	log.Printf("Starting server on %s\n", cfg.API.Port)
	log.Fatal(http.ListenAndServe(cfg.API.Port, server))
}

func sendError(err error) {
	log.Fatalf("Error: %s", err)
	os.Exit(1)
}

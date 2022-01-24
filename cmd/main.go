package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom/config"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	api "github.com/Vivi3008/apiTestGolang/http"
	"github.com/Vivi3008/apiTestGolang/store"
	_ "github.com/lib/pq"
)

func main() {
	/* 	err := godotenv.Load(".env")

	   	if err != nil {
	   		log.Fatalf("Error loading .env file: %s", err)
	   	} */

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load configuration")
	}

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	addr := ":3000"

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

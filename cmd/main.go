package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	api "github.com/Vivi3008/apiTestGolang/http"
	"github.com/Vivi3008/apiTestGolang/store"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Connect to database sucessfully'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

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

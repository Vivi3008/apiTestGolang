package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/commom/config"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/postgres"

	acStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/account"
	atStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/activity"
	blStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/bills"
	trStore "github.com/Vivi3008/apiTestGolang/gateways/db/store/transfers"

	account_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/account"
	activities_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/activity"
	bills_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/bills"
	transfers_postgres "github.com/Vivi3008/apiTestGolang/gateways/db/postgres/entries/transfers"

	api "github.com/Vivi3008/apiTestGolang/gateways/http"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
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

	var accountUsecase account.AccountUsecase
	var transferUsecase transfers.TranfersUsecase
	var billsUsecase bill.BillUsecase
	var activityUsecase activities.ActivityUsecase

	_, ok := os.LookupEnv("LOCAL_STORAGE")

	if ok {
		accountStore := acStore.NewAccountStore()
		transStore := trStore.NewTransferStore()
		billStore := blStore.NewBillStore()
		activitiesStore := atStore.NewAccountActivity()

		accountUsecase = account.NewAccountUsecase(accountStore)
		transferUsecase = transfers.NewTransferUsecase(transStore, accountUsecase)
		billsUsecase = bill.NewBillUseCase(billStore, accountUsecase)
		activityUsecase = activities.NewAccountActivityUsecase(activitiesStore)
	} else {
		accountStore := account_postgres.NewRepository(db)
		transStore := transfers_postgres.NewRepository(db)
		billStore := bills_postgres.NewRepository(db)
		activitiesStore := activities_postgres.NewRepository(db)

		accountUsecase = account.NewAccountUsecase(accountStore)
		transferUsecase = transfers.NewTransferUsecase(transStore, accountUsecase)
		billsUsecase = bill.NewBillUseCase(billStore, accountUsecase)
		activityUsecase = activities.NewAccountActivityUsecase(activitiesStore)
	}

	server := api.NewServer(accountUsecase, transferUsecase, billsUsecase, activityUsecase)

	log.Printf("Starting server on %s\n", cfg.API.Port)
	log.Fatal(http.ListenAndServe(cfg.API.Port, server))
}

func sendError(err error) {
	log.Fatalf("Error: %s", err)
	os.Exit(1)
}

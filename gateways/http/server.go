package http

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/http/accounts"
	"github.com/Vivi3008/apiTestGolang/gateways/http/auth"
	bl "github.com/Vivi3008/apiTestGolang/gateways/http/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	transfer "github.com/Vivi3008/apiTestGolang/gateways/http/transfers"

	"github.com/gorilla/mux"
)

func NewServer(
	accountUc account.AccountUsecase,
	usecaseTr transfers.TranfersUsecase,
	usecaseBl bill.BillUsecase,
) http.Handler {
	router := mux.NewRouter()
	routerAuth := router.NewRoute().Subrouter()

	accounts.NewHandler(router, accountUc)
	auth.NewHandler(router, accountUc)
	transfer.NewHandler(routerAuth, usecaseTr, accountUc)
	bl.NewHandler(routerAuth, usecaseBl, accountUc)

	routerAuth.Use(middlewares.Auth)

	return router
}

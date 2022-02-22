package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

var (
	ErrAuth = errors.New("authentication required")
)

type AuthContextKey string

const HeaderKey = "Authorization"

var ContextAccountID = AuthContextKey("account_id")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header[HeaderKey] == nil {
			response.SendError(w, ErrAuth, http.StatusUnauthorized)
			return
		}

		authHeader := r.Header.Get(HeaderKey)

		//pegar o id do token
		accountId, err := commom.AuthJwt(authHeader)

		if err != nil {
			response.SendError(w, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextAccountID, accountId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAccountId(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(ContextAccountID).(string)
	return tokenStr, ok
}

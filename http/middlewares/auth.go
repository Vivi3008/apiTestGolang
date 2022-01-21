package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/http/response"
)

var (
	ErrAuth = errors.New("authentication required")
)

type AuthContextKey string

var contextAccountID = AuthContextKey("account_id")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			response.SendError(w, ErrAuth, http.StatusUnauthorized)
			return
		}

		authHeader := r.Header.Get("Authorization")

		//pegar o id do token
		accountId, err := commom.AuthJwt(authHeader)

		if err != nil {
			response.SendError(w, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextAccountID, accountId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAccountId(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(contextAccountID).(string)
	return tokenStr, ok
}

package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom"
)

var (
	ErrAuth = errors.New("authentication required")
)

type AuthContextKey string

var contextAccountID = AuthContextKey("account_id")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Authentication required")

		}

		authHeader := r.Header.Get("Authorization")

		//pegar o id do token
		accountId, err := commom.AuthJwt(authHeader)

		if err != nil {
			log.Printf("Login failed: %s\n", err.Error())
			response := Error{Reason: err.Error()}
			w.Header().Set(ContentType, JSONContentType)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
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

package commom

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthJwt(tokenStr string) (string, error) {
	var accountId string

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountId = claims["id"].(string)
	} else {
		return "", err
	}

	return accountId, nil
}

package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"rides/models"
	"rides/util"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		token := ""
		if _, ok := req.Header["Authorization"]; ok && strings.HasPrefix(req.Header.Get("Authorization"), "Bearer") {
			token = strings.Fields(req.Header.Get("Authorization"))[1]
		}

		if token == "" {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		keyFunc := func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Invalid Token")
			}
			return []byte(util.AccessToken()), nil
		}

		jwtToken, err := jwt.ParseWithClaims(token, &models.OrganizationClaims{}, keyFunc)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				res.WriteHeader(http.StatusUnauthorized)
				return
			}
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if !jwtToken.Valid {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		orgClaims, ok := jwtToken.Claims.(*models.OrganizationClaims)
		if !ok {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println(orgClaims)

		next(res, req)
	}
}

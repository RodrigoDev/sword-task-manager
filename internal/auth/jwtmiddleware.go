package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type ErrorResponse struct {
	msg string
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(ErrorResponse{msg: "missing auth token"})
			return
		}
		token := &jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(ErrorResponse{msg: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "props", *token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

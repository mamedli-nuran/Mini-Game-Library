package middleware

import (
	"context"
	"mini-game-library/constant"
	"mini-game-library/handler"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JWTMiddleware(jwtSecret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		scheme := "Bearer "
		if len(authHeader) < len(scheme) {
			handler.WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		userScheme := authHeader[:len(scheme)]
		if !strings.EqualFold(scheme, userScheme) {
			handler.WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		userJWT := authHeader[len(scheme):]
		var claims jwt.RegisteredClaims
		_, err := jwt.ParseWithClaims(
			userJWT,
			&claims,
			func(t *jwt.Token) (any, error) {
				return []byte(jwtSecret), nil
			},
			jwt.WithValidMethods([]string{"HS256"}),
		)

		if err != nil {
			handler.WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			handler.WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		}

		ctx := context.WithValue(r.Context(), constant.UserIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}

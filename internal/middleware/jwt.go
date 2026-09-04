package middleware

import (
	"context"
	constant2 "mini-game-library/internal/constant"
	"mini-game-library/internal/handler"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JWTMiddleware(jwtSecret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.WriteError(w, r, http.StatusUnauthorized, constant2.ErrUnauthorized)
			return
		}

		scheme := "Bearer "
		if len(authHeader) < len(scheme) {
			handler.WriteError(w, r, http.StatusUnauthorized, constant2.ErrUnauthorized)
			return
		}

		userScheme := authHeader[:len(scheme)]
		if !strings.EqualFold(scheme, userScheme) {
			handler.WriteError(w, r, http.StatusUnauthorized, constant2.ErrUnauthorized)
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
			handler.WriteError(w, r, http.StatusUnauthorized, constant2.ErrUnauthorized)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			handler.WriteError(w, r, http.StatusUnauthorized, constant2.ErrUnauthorized)
		}

		ctx := context.WithValue(r.Context(), constant2.UserIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}

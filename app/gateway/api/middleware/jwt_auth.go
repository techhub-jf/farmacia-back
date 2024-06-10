package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
)

func ProtectedHandler(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")

			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Missing authorization header")

				return
			}

			tokenString = tokenString[len("Bearer "):]

			claims, err := verifyToken(tokenString, jwtSecret)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid token")

				return
			}

			if claims != nil {
				if r.Context().Value(dto.User{}) == nil {
					user := dto.User{
						ID: claims["user"],
					}
					r = r.WithContext(context.WithValue(r.Context(), dto.User{}, user))
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func verifyToken(tokenString string, jwtSecret string) (jwt.MapClaims, error) {
	secretKey := []byte(jwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //nolint:revive
		return secretKey, nil
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims) //nolint:forcetypeassert

	return claims, nil
}

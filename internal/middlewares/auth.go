package middlewares

import (
	"context"
	"github.com/dshum/school/internal/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		} else {
			token, _ := utils.ParseToken(authHeader[1])

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "auth", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("auth").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}
		}
	})
}

package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/analopesdev/duochat-service/internal/auth"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("uid")

			if err != nil || c.Value == "" {
				json.NewEncoder(w).Encode(map[string]string{"message": "unauthorized"})
				return
			}

			claims, err := auth.ParseToken(c.Value)

			if err != nil || claims == nil {
				json.NewEncoder(w).Encode(map[string]string{"message": "invalid token"})
				return
			}

			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"message": "invalid token"})
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

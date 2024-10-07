package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			err := godotenv.Load()
			if err != nil {
				fmt.Println("Error loading .env file")
			}
			jwtKey := os.Getenv("JWT_SECRET")
			return []byte(jwtKey), nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
		jwtKey := os.Getenv("JWT_SECRET")
		return []byte(jwtKey), nil
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"message":  "Token is valid",
		"username": claims.Username,
		"role":     claims.Role,
		"id":       strconv.Itoa(claims.Id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetClaimsFromContext(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value("claims").(*Claims)
	return claims, ok
}

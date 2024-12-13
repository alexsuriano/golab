package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-123")

const webServerPort = ":8585"

func main() {
	fmt.Printf("Servidor rodando na porta %v...", webServerPort)

	http.HandleFunc("/protected", ProtectedEndpoint)
	http.ListenAndServe(webServerPort, nil)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v\n", token.Header["alg"])
		}
		return jwtSecret, nil
	})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		http.Error(w, "Token não encontrado", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Acesso autorizado ao endpoint protegido"))
}

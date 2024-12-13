package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-123")

const webServerURL = "http://localhost:8585/protected"

func main() {
	token, err := CreateJwt()
	if err != nil {
		log.Println(err)
	}

	log.Println(token)

	SendRequest(token)
}

func CreateJwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "client123",                            //ID do cliente
		"iat": time.Now().Unix(),                      //Timestamp de criacao
		"exp": time.Now().Add(time.Minute * 5).Unix(), //Expiracao do token em 5min
	})

	//gera uma string assinada com o segredo
	return token.SignedString(jwtSecret)
}

func SendRequest(token string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", webServerURL, nil)
	if err != nil {
		log.Println("Erro ao criar requisição")
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Acesso permitido! Http status %v\n", resp.StatusCode)
	} else {
		fmt.Printf("Acesso negado! Http status %v\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Body da requisição: %v\n", string(body))
}

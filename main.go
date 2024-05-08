package main

import (
	"fmt"
	"os"

	clients "github.com/RubenLL/GOCognitoTest/Clients"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	CLOUD_REGION := os.Getenv("CLOUD_REGION")
	COGNITO_CLIENT := os.Getenv("COGNITO_CLIENT")

	client := clients.NewCognitoClient(CLOUD_REGION, COGNITO_CLIENT)
	fmt.Println(client.AppClientID)
	/*
		result, err := client.SignUp("teste", "passW0rdD3Teste@", "ruben.lopez.deleon@gmail.com")

		if err != nil {
			panic(err)
		}
		fmt.Printf("cadastro realizado %v\n", result)
	*/

	/*

			result, err := client.ConfirmSignUp("ruben.lopez.deleon@gmail.com", "099962")

			if err != nil {
				panic(err)
			}
			fmt.Printf("Confirmacão realizada %v\n", result)
		}
	*/
	result, JWT, err := client.SignIn("ruben.lopez.deleon@gmail.com", "passW0rdD3Teste@")

	if err != nil {
		panic(err)
	}
	fmt.Printf("Login realizado %v\n", result)
	fmt.Printf("JWT: %v\n", JWT)
}

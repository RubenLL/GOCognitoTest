package main

import (
	"fmt"

	clients "github.com/RubenLL/GOCognitoTest/Clients"
)

func main() {
	fmt.Println("Hola mundo")
	client := clients.NewCognitoClient("sa-east-1", "sa-east-1_lErE6Q9E4")
	fmt.Println(client.AppClientID)
	result, err := client.SignUp("teste", "passW0rdD3Teste@", "ruben.lopez.deleon@gmail.com")

	if err != nil {
		panic(err)
	}
	fmt.Printf("cadastro realizado %v\n", result)
}

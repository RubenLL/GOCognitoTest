package clients

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var secretHash = "dh569vqr4vidbbk0a9pc79no93qe5o223c3q14j97vk5fukiv0g"

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
	AppClientID   string
}

func (actor CognitoActions) SignUp(userName string, password string, userEmail string) (bool, error) {
	confirmed := false

	output, err := actor.CognitoClient.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.AppClientID),
		Password: aws.String(password),
		Username: aws.String(userEmail),
		//SecretHash: aws.String(secretHash),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(userEmail)},
			//{Name: aws.String("name.formated"), Value: aws.String(userName)},
		},
	})
	fmt.Printf("ActorData: %v", actor)
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't sign up user %v. Here's why: %v\n", userName, err)
		}
	} else {
		confirmed = output.UserConfirmed
	}
	return confirmed, err
}

func NewCognitoClient(region string, cognitoID string) CognitoActions {
	conf, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region))
	if err != nil {
		panic(err)
	}
	cognitoClient := cognitoidentityprovider.NewFromConfig(conf)

	cognito := CognitoActions{
		CognitoClient: cognitoClient,
		AppClientID:   cognitoID,
	}
	return cognito

}

func (actor CognitoActions) ConfirmSignUp(email string, code string) (bool, error) {

	confirmed := false
	output, err := actor.CognitoClient.ConfirmSignUp(context.TODO(), &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(actor.AppClientID),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	})

	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println(output)
		return confirmed, err

	}

	confirmed = true
	fmt.Println(output)
	return confirmed, err
}

func (actor CognitoActions) SignIn(userID string, password string) (bool, string, error) {
	output, err := actor.CognitoClient.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(actor.AppClientID),
		AuthParameters: map[string]string{"USERNAME": userID, "PASSWORD": password},
	})
	if err != nil {
		return false, "", err
	}
	return true, *output.AuthenticationResult.AccessToken, nil
}

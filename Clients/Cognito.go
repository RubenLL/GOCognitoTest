package clients

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
	AppClientID   string
}

func (actor CognitoActions) SignUp(userName string, password string, userEmail string) (bool, error) {
	confirmed := false
	output, err := actor.CognitoClient.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.AppClientID),
		Password: aws.String(password),
		Username: aws.String(userName),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(userEmail)},
		},
	})
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

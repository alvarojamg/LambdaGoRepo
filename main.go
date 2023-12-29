package main

import (
	"context"
	"errors"
	"fmt"
	"main/awsgo"
	"main/config"
	"main/models"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(EjecutLambda)

}

func EjecutLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAws()

	if !ValidateParams() {
		fmt.Println("Error in params. send SecretManager")
		err := errors.New("Error in params, most send SecretManager")
		return event, err
	}
	var data models.SiginUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("email=" + data.UserEmail)
		case "sub":
			data.UserUUId = att
			fmt.Println("Sub" + data.UserUUId)

		}
	}

	err := config.ReadSecret()
	if err != nil {
		fmt.Println("Error at reading secret=" + err.Error())
		return event, err
	}

	err = config.SignUp(data)
	return event, err

}

func ValidateParams() bool {
	var getParams bool
	_, getParams = os.LookupEnv("SecretName")
	return getParams
}

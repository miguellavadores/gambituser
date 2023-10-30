package main

import (
	"context"
	"fmt"
	"os"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/miguellavadores/gambituser/awsgo"
	"github.com/miguellavadores/gambituser/models"
	"github.com/miguellavadores/gambituser/bd"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	
	awsgo.InicializoAWS()
	
	if !ValidoParametros() {
		fmt.Println("Error en los parametros. Debe enviar 'SecretName'")
		err := errors.New("Error en los paráametros. Debe enviar 'SecretName'")
		return event, err
	}
	
	var datos models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = "+ datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = "+ datos.UserUUID)
		}
	}
	
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret"+err.Error())
		return event, err
	}
	
	err = bd.SignUp(datos)
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
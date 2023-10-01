package awsgo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-2"))
	
	if err != nil {
		panic("Error al cargar la configurations .aws/config "+err.Error())
	}
	
}
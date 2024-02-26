package main

import(
	"fmt"
	"os"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"context"
)

func main() {
	lambda.Start(EjecutarLambda)
}

func EjecutarLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

}

func ValidarParametros()bool{
	_, traerParametro := os.LookupEnv("SecretName")
	if !traerParametro{
		return traerParametro
	}

	_, traerParametro = os.LookupEnv("BucketName")
	if !traerParametro{
		return traerParametro
	}

	_, traerParametro = os.LookupEnv("UrlPrefix")
	if !traerParametro{
		return traerParametro
	}
}
package main

import (
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/awsgo"
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/bd"
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/handlers"
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/models"
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/secretmanager"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutarLambda)
}

func EjecutarLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !ValidarParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. deben incluir 'SecretName', 'BucketName', 'UrlPrefix' ",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la Lectura del 'Secret'" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)
	fmt.Println("Path:" + path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketname"), os.Getenv("BucketName"))

	fmt.Println("user:" + SecretModel.Username + " Password: " + SecretModel.Password + " Host: " + SecretModel.Host)

	// Chequeo Conexión a la BD o Conecto la BD

	bd.ConectarDB(awsgo.Ctx)

	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	fmt.Println("Sali de Manejadores")

	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}

}

func ValidarParametros() bool {
	_, traerParametro := os.LookupEnv("SecretName")
	if !traerParametro {
		return traerParametro
	}

	_, traerParametro = os.LookupEnv("BucketName")
	if !traerParametro {
		return traerParametro
	}

	_, traerParametro = os.LookupEnv("UrlPrefix")
	if !traerParametro {
		return traerParametro
	}

	return traerParametro
}

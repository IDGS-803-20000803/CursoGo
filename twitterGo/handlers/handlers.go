package handlers

import (
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/models"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + "> " + ctx.Value(models.Key("method")).(string))
	var r models.RespAPI
	r.Status = 400
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Method Invalid"
	return r

}

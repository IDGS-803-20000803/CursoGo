package bd

import (
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectarDB(ctx context.Context) error {
	
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=Twitter", user, password, host)

	log.Println("conexion bd = " + connStr)

	opts := options.Client().ApplyURI("mongodb+srv://root:Poder335@twitter.qqi28kf.mongodb.net/?retryWrites=true&w=majority&appName=Twitter")
	//var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Realizando Ping")

	/*
			if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
				fmt.Println(err.Error())
				return err
			}
			fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")


		err = client.Database(DatabaseName).Client().Ping(context.TODO(), nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	*/

	fmt.Println("Conexion exitosa con la BD")
	MongoCN = client
	db := ctx.Value(models.Key("database")).(string)
	DatabaseName = string(db)
	return nil

}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}

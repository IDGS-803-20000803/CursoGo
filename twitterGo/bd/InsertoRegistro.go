package bd

import (
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	// Verificar si la conexión a MongoDB está establecida
	if MongoCN == nil {
		fmt.Println("No se ha establecido conexión a la base de datos MongoDB")
	}

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	// Encriptar la contraseña
	hashedPassword, err := EncriptarPassword(u.Password)
	if err != nil {
		fmt.Println("Error al encriptar la contraseña")
	}

	// Actualizar el campo de contraseña con el valor encriptado
	u.Password = hashedPassword

	fmt.Println("Usuario:", u)

	result, err := col.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println("Error al insertar el usuario en la base de datos:", err)
		return "", false, err
	} else {
		fmt.Println("Usuario insertado correctamente")
	}

	ObjID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("Error al obtener el ID del objeto insertado")
	}

	return ObjID.String(), true, nil
}

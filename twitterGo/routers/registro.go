package routers

import (
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/bd"
	"C/Users/sistemas8/Documents/cursoGo/twitterGo/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func Registro(ctx context.Context) models.RespAPI {
	var t models.Usuario
	var r models.RespAPI
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)

	fmt.Println("Body:" + body)
	//err := json.Unmarshal([]byte(body), &t)
	err := json.NewDecoder(strings.NewReader(body)).Decode(&t)

	if err != nil {
		r.Message = "Error al decodificar JSON: " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	fmt.Println("Usuario después de la deserialización:", t)

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Println(r.Message)
		return r
	}

	fmt.Println("Paso la validacion de Email", t.Email)

	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}
	fmt.Println("Paso la validacion de Password", t.Password)

	fmt.Println("Entrando a funcion de si el usuario ya existe")
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	fmt.Println("Saliendo de validacion de existencia de usuario")

	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese Email"
		fmt.Println(r.Message)
		return r
	}
	fmt.Println("Entrando a funcion de insertar registro")
	_, status, err := bd.InsertoRegistro(t)
	fmt.Println("Saliendo de la funcion de insertar registro")
	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r
}

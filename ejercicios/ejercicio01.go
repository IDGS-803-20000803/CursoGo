package ejercicios

import (
	"strconv"
	"log"
)

func ConverirtEntero(numero string)( int,  string){
	var num int
	var cadena string
	var err error
	
	num, err = strconv.Atoi(numero)

	if err != nil {
		log.Fatal(err)
	}
	
	if num > 100 {
		cadena = "El numero es mayor a 100"
	} else {
		cadena = "El numero es menor a 100"
	}

	return num, cadena
}
package ejercicios

import (
	"fmt"
)

//La funcion Fprintf muestra en consola
//La funcion Sprintf devuelve un string

func TablaMultiplicar(numero int) string {
	var cadena string
	for i := 1; i <= 10; i++ {
       cadena += fmt.Sprintf("%d x %d = %d\n", numero, i, numero * i)
    }
	return cadena
}
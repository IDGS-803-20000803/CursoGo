package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

//Estandar Input por teclado Stdin
//Estandar OutPut por Pantalla
//Todo lo que entra por Bufio entra como texto

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese un numero: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Print("Ingrese otro numero: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Print("Ingrese la leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
	/*fmt.Scan(&numero1)
	  fmt.Print("Ingrese otro numero: ")
	  fmt.Scan(&numero2)*/
}

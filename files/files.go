package files

import (
	"C/Users/sistemas8/Documents/cursoGo/ejercicios"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var fielName string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var numero1 int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese el numero de la tabla de multiplicar a realizar: ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
		}
		if err != nil {
			fmt.Println("Error al capturar información" + err.Error())
			continue
		} else {
			texto = fmt.Sprintln(ejercicios.TablaMultiplicar(numero1))
			archivo, err := os.Create(fielName)
			if err != nil {
				fmt.Println("Error al crear el archivo" + err.Error())
				return
			}
			fmt.Fprintf(archivo, texto)
			archivo.Close()
			break
		}

	}

}

func SumarTabla() {
	var numero1 int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese el numero de la tabla de multiplicar a realizar: ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
		}
		if err != nil {
			fmt.Println("Error al capturar información" + err.Error())
			continue
		} else {
			texto = fmt.Sprintln(ejercicios.TablaMultiplicar(numero1))
			if Append(fielName, texto) == false {
				fmt.Println("Error al escribir en el archivo")
			}
			break
		}

	}

}

func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fielName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

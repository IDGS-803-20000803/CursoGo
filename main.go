package main

import (
	//"C/Users/sistemas8/Documents/cursoGo/goroutines"
	//"C/Users/sistemas8/Documents/cursoGo/ejercicios"
	//"C/Users/sistemas8/Documents/cursoGo/variables"
)

/* =========Condicionales============================
if os := runtime.GOOS; os == "Linux."{
	fmt.Println("Esto no es Windows")
}else {
	fmt.Println("Esto es Windows")
}

switch os := runtime.GOOS; os{
case "linux":
	fmt.Println("Esto es Linux")

case "windows":
	fmt.Println("Esto es Windows")
default:
	fmt.Printf("%s \n", os)
}
*/

/*
	==========Conversion cadena a Entero y viceversa===============

num1, texto := ejercicios.ConverirtEntero("450")
fmt.Println(num1)
fmt.Println(texto)
estado, texto :=	variables.ConvertirTexto(1555)
fmt.Println("Estado =",estado)
fmt.Println("texto =",texto)
*/

/* =========Ingreso por teclado=================
teclado.IngresoNumeros()
*/

/* ========Ciclo FOR=========================

for i := 0; i < 10; i++{
		if i == 8 {
			break
		}
        fmt.Println(i)
    }

scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Ingrese el numero de la tabla de multiplicar a realizar: ")
			if scanner.Scan() {
				numero1, err = strconv.Atoi(scanner.Text())
			}
			if err!= nil {
				fmt.Print("Error al capturar información" + err.Error())
	            continue
	        } else {
				fmt.Println(ejercicios.TablaMultiplicar(numero1))
				break
			}

		}

iteraciones.Iterar()
*/

/* ==========Leer y Escribir en archivo.txt================

files.LeoArchivo()
files.GrabarTabla()

*/

/* ==========Funciones Anonimas, Closures ===================
arreglos_slices.Capacidad()
*/

/* ===========Mapeo: Clave Valor
mapas.MostrarMapas()
*/
/* ===Modelos===========
users.AltaUsuario()
*/

/*======Interfaces====================
Pedro := new(modelos.Hombre)
e.HumanosRespirando(Pedro
Maria := new(modelos.Mujer)
e.HumanosRespirando(Maria)
*/

/* =========Defer, panic y recover=====
defer_panic.EjemploPanic()
*/

/* ========Go Rutines ======================
	canal1 := make(chan bool)
	go goroutines.MinombreLentoo("Daniel Sanchez", canal1)
	fmt.Println("Estoy aqui")

	defer func() {
		<-canal1
	}()
 */

func main() {
	/*
		middleware.MiMiddleware()
		
		webserver.MiwebServer()
	*/
	

}

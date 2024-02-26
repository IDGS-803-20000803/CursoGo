package middleware

import "fmt"

func sumar(a,b int) int {
	return a + b
}

func restar(a,b int) int {
	return a - b
}

func multiplicar(a,b int) int {
	return a * b
}

func MiMiddleware(){
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2,3)
	fmt.Println(result)
	
    result = operacionesMidd(restar)(2,3)
	fmt.Println(result)
	
    result = operacionesMidd(multiplicar)(2,3)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
    return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
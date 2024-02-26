package ejer_interfaces

import (
	"C/Users/sistemas8/Documents/cursoGo/interfaces"
	"fmt"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando,", hu.Sexo())
}

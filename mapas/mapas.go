package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogotá"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Argentina": 4,
		"Brasil":    2,
		"Uruguay":   3,
		"Chile":     1,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo: %s, tiene un puntaje: %d", equipo, puntaje)
	}

	delete(campeonato, "Brasil")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("Puntaje: %d y el equipo existe = %t", puntaje, existe)

}

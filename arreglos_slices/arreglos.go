package arreglos_slices

import "fmt"

var tabla [10] int = [10]int {10,0,59}
var matriz [5][6]int

func MuestroArreglos(){
	tabla[7] = 33
	tabla[8] = 4

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

//Recorremos arreglos
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
    
}

   /* for i := 0; i < 5; i++ {
        for j := 0; j < 6; j++ {
            fmt.Println(matriz[i][j])
        }
    }*/
	
}
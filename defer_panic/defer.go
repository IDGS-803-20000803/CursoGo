package defer_panic

import (
	"log"
	"fmt"
)

func VemosDefer(){
	//El defer es la accion que realizara hasta el final de la funcion inlcuso si ocurre un error 
	//en tiempo de ejecucion
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final de la funcion")
	fmt.Println("Este es el segundo mensaje")
}
	//El Panic detiene la aplicacion mostrando un mensaje
	//El Recover permite recuperar el panic para guarda el error en logs
func EjemploPanic(){

	defer func() {
        reco := recover()
		if reco!= nil {
            log.Fatalf("Ocurrio un error que genero Panic: %v", reco)
        }
    }()
	a:= 1
	if a== 1{
		panic("Se encontro el valor 1")
	}
}
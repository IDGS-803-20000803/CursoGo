package funciones

import "fmt"

func tabla(valor int) func() int{
	numero := valor
	secuencia := 0
	return func() int{
        secuencia++
        return numero*secuencia
    }
}

func LlamarClosures(){
	tabladel := 2
	miTabla := tabla(tabladel)
	for i := 0; i < 10; i++ {
		fmt.Println(miTabla())
	}
}
package iteraciones

import (
	"fmt"
)

func Iterar(){
	for i := 0; i < 10; i++{
		if i == 8 {
			break
		}
        fmt.Println(i)
    }
	
}
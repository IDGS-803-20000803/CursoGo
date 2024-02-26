package users

import (
	"C/Users/sistemas8/Documents/cursoGo/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Juan", time.Now(), true)
	fmt.Println(u)
}

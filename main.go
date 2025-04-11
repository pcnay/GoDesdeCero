package main

import (
	"github.com/DevOps-PcNay/GoDesdeCero/Variables"
	"fmt"
)

func main() {	
	//Variables.MuestroEnteros()
	//Variables.RestoVariables()
	estado,cadena := Variables.ConviertoaTexto(652145)
	fmt.Println("Estado ",estado," Contenido ",cadena)

}

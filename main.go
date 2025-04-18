package main

import (
	"github.com/DevOps-PcNay/GoDesdeCero/Variables"
	"github.com/DevOps-PcNay/GoDesdeCero/Ejercicios"
	"fmt"	
)

func main() {	
	//Variables.MuestroEnteros()
	//Variables.RestoVariables()
	estado,cadena := Variables.ConviertoaTexto(652145)
	fmt.Println("Estado ",estado," Contenido ",cadena)
	//Variables.Condicionales_If()
	Variables.Switch()
	
	return_number,messajes := Ejercicios.Recibe("sdfsfdsfd")
	fmt.Printf("Valor enviado %d , Mensaje : %s \n",return_number,messajes)

	

}

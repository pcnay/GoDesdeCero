package Funciones

import (
	"fmt"
)

func Calculos() {
	//var numero1 int
	//var numero2,numero5 int

	// Asignacion de funcion anonima.
	calculo := func(numero3 int, numero4 int) (resultado int) {
		return numero3 + numero4
	}
	fmt.Println(calculo(20, 34))

	// Solo se coloca "=" porque arriba se define con el simbolo ":="
	// Se reutiliza la funcion anomina definida arriba, solo se debe cambiar la logica.
	calculo = func(numero1 int, numero2 int) (resultado2 int) {
		return numero1 * numero2
	}

	fmt.Println(calculo(340, 203))
}

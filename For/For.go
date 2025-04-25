package For

import (
	"fmt"
)

// Se utilizan en los Arreglos, Slices, Maps, "rows' de consultas de de datos.  para recorrer el contenido.
// Es la unica instruccion para recorrer elementos

func For_ciclado() {

	// For ciclado, funcionamiento similar al While
	for {
		fmt.Println("hola")
		break
	}
}

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println("Mostrando valor de i = ", i)
		i++
	}
}

func Iterar2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Mostrando valor de i (usando for simplificado )", i)
	}

}

// Iterando con la instruccion "range"
func Iterar_Range() {
	for l := range 10 {
		fmt.Println("Mostrao el valor de l (usando la funcion Range) ", l)
	}
}

// For decrementando
func Iterar_Reverse() {
	for k := 100; k > 10; k -= 5 {
		fmt.Println("Decrementando el Valor ", k)
	}
}

// For con pase a la siguiente instruccion
func Iterar_ForContinue() {
	for m := 20; m > 3; m -= 2 {
		if m == 10 {
			continue // No imprime el numero 10
		}
		fmt.Println("Imprimiendo el contador en decremento usando - Continue -", m)
	}
}

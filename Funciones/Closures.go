package Funciones

import (
	"fmt"
)

func tabla(valor int) func() (resultado int) {
	numero := valor
	secuencia := 0

	// Este es la parte que se utiliza para Closure
	return func() (resultado int) {
		secuencia++
		return numero * secuencia
	}

}

func LlamarClosure() {
	var tabladel int = 2
	tablaMultiplicar := tabla(tabladel)
	for k := 1; k <= 10; k++ {
		fmt.Println(tablaMultiplicar())
	}

}

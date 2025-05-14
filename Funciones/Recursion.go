package Funciones

import (
	"fmt"
)

func Exponencial(valor int) {
	var valor_mostrado string
	if valor > 10000 {
		return
	}

	valor_mostrado = fmt.Sprintf("Contenido de Valor = %d \n", valor)
	fmt.Println(valor_mostrado)

	Exponencial(valor * 2)

}

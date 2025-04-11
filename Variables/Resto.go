package Variables

import (
	"fmt"
	"time"
	"strconv"
)

// Se definen las variables de forma global, se puede ver desde otros archivos
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (estado bool, texto string) {
	texto = strconv.Itoa(numero)
	return true,texto

}

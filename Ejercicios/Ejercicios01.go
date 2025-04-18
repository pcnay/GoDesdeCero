package Ejercicios
import (
	"fmt"
	"strconv"

)

func Recibe (number_strings string ) (valor int, valor2 string){
	// Convertir el valor recibo (numero ) a Entero
	// EL paquete se llama "strconv" la funcion "Atoi"
	// "Atoi", retorna dos valores, el numero convertido y un error
	var cadena_num string 

	num_conv,err := strconv.Atoi(number_strings)
	if err != nil{
		fmt.Printf("Error Found %T \n",err.Error())
		return 0,"error"
	}

	if num_conv > 100 {
		cadena_num = "El numero es mayor a 100"
		}

	if num_conv < 100 {
		cadena_num = "El numero es Menor a 100"
	}
return num_conv,cadena_num
}
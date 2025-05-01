package Ejercicios

/*
1.- Crear un archivo en Go en el paquete Ejercicios, llamado "Ejercicio02.go"
2.- Crear una funcion publica que pida por teclado un numero y valide si hay error o no (si hay error que vuelva a pedirlo)
3.-En base al numero recibido crear una tabla numerica del mismo. De 1 al 10 y la muestre en pantalla.
4.-En Main llamar a dicha funcion.
5.- Grabar, ejecutar y subirlo a Github.com
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dato int
var err error
var text string // Se modificara para grabar la tabla de numeros en un archivo.

// func Ingresar_numero(numero int) = Se modifica la funcion para que retorne la tabla de la multiplicacion
func Ingresar_numero() (contenido_tabla string) {
	var contador int = 0
	for {
		fmt.Println("Ingresa el numero -1 = Salir  : ")
		teclado := bufio.NewScanner(os.Stdin)

		if teclado.Scan() { // Si tecleo un dato el usuario
			dato, err = strconv.Atoi(teclado.Text())
		}
		if err != nil {
			fmt.Printf(" Dato Incorrecto \n ")
			//panic("El dato ingresado es incorrecto " + err.Error())
		}
		if dato == -1 {
			break
		}
		if dato != 0 {
			contador++

			if contador > 1 {
				text += fmt.Sprintf("\n \n")
			}

			for k := 1; k <= 10; k++ {
				// Se modificara para que grabe en un archivo.
				//fmt.Printf(" %d X %d = %d \n", dato, k, (dato * k))
				// Se utilizara otra instruccion, "Sprintf" para que imprima una cadena la salida del "Printf"

				text += fmt.Sprintf(" %d x %d = %d \n", dato, k, (dato * k))
			} // for k := 0; k <= 0...

		} //if dato != 0 {

	} // For
	return text
}

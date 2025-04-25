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
var salir bool = true

func Ingresar_numero(numero int) {

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
			for k := 1; k <= 10; k++ {
				fmt.Printf(" %d X %d = %d \n", dato, k, (dato * k))
			} // if dato != 0

		}

		//fmt.Println(dato)

	} // For

}

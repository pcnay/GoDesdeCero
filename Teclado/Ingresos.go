package Teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese numero 1 :")
	// Pidiendo por teclado.
	// Se puede leer desde varias fuente de datos .
	scanner := bufio.NewScanner(os.Stdin) // Teclado

	if scanner.Scan() { // Si tecleo un dato el usuario
		// No se utiliza asignacion directa, ya que la variable se declara en la parte de arriba

		numero1, err = strconv.Atoi(scanner.Text()) // Conv. de Cadena a Entero
		if err != nil {
			// Ya no se ejecutara, se detiene
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese numer 2 :")
	scanner2 := bufio.NewScanner(os.Stdin)
	if scanner2.Scan() {
		numero2, err = strconv.Atoi(scanner2.Text())
		if err != nil {
			panic("El dato No. 2 ingresado es incorrecto " + err.Error())
		}

	}

	fmt.Println("Ingrese leyenda :")

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

} // if scanner.Scan() { // Si tecleo un dato el usuario

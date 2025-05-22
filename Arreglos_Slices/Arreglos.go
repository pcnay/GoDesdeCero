package Arreglos_Slices

import (
	"fmt"
)

var tabla [10]int

// Otra forma de declarar los arreglos.
var tabla2 [20]int = [20]int{20, 23, 25, 29, 21, 0, 0, 0, -1, 23, 33, 40, 50}

// Definiendo un arreglo de cadenas
var tablaCadenas [10]string = [10]string{"Juan", "Pedro", "Pablo"}

// Definiendo una matriz bidireccional.
var matriz [20][30]int

func MuestroArreglos() {
	//Asignacion directa para valores de Arreglos.
	tabla[7] = 33
	tabla[2] = 10
	fmt.Println("Mostrando Valores del Arreglo \n", tabla)

	fmt.Println("Mostrando el contenido del arreglo 2 \n", tabla2)
	fmt.Println("Mostrando el contenido del arreglo de Cadenas \n", tablaCadenas)

	fmt.Println("Recorriendo un arreglo")

	for k := 1; k < len(tabla); k++ {
		fmt.Println("Desplegando el contenido del arreglo \n", tabla[k])
	}

	// Asignando un valor a la matriz
	matriz[1][1] = 30
	fmt.Println(matriz)

} // func MuestroArreglos() {

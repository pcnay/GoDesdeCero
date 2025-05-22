package Arreglos_Slices

import (
	"fmt"
)

// Slices son vectores pero no se le asignan una capacidad inicial
var tablaS []int = []int{2, 4, 6}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	// Tener encuenta que no se espeficica el largo de la dimension del vector, que es una particularidad de los Slices.
	// Crendo el Slice, desde el elemento 3,  (iniciando desde 0 se cuentan 3 posiciones "67" hasta el 9)
	porcion := arreglo[3:] // Se crea e inicializa la variable
	fmt.Println(porcion)
	// Creando otros Slices desde la posicion 0 hasta la 5 son posicion cardinales (el 0 es una posicion)
	porcion2 := arreglo[:5]
	fmt.Println(porcion2)

	// Definiendo un Slices de forma que se extraiga del vector
	porcion3 := arreglo[6:7]
	fmt.Println(porcion3)
}

func Capacidad() {
	// Se crea un Slices incluyendo su capacidad.
	// Con esta forma de asignar la capacidad, el desempeno es superior, ya que reserva por adelantado la memoria
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo = %d, Capacidad = %d \n", len(elementos), cap(elementos))

	// Creando un Slices vacio, para posteriormente agregar elementos.
	nums := make([]int, 0, 0)
	for k := 0; k <= 100; k++ {
		nums = append(nums, k)
	}

	// Se autoajusta la capacidad de forma 2 ^ n, se asigna el espacio en memoria.
	// Ya que es mas costoso reservar memoria y solo utilizarla (previamente reservada cuando se crea el Slices)
	// Se usan en calculos que los arreglos crecen aleatoriamente, por lo que es muy utilizado en cuestiones cientificas.
	fmt.Printf("Largo = %d, Capacidad = %d \n", len(nums), cap(nums))
}

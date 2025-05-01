package Files

import (
	"bufio"
	"fmt"
	"github.com/DevOps-PcNay/GoDesdeCero/Ejercicios"
	"os"
)

var filename string = "./Files/Txt/Tabla.txt"

// Graba la tabka en un archivo
func Grabar_tabla() {
	var texto string = Ejercicios.Ingresar_numero()

	// Crea el archivo, pero si existe lo borra, por esta razon ses crea la funcion "Suma_tabla"
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error Al Crear El Archivo " + err.Error())
		return // Sale de la funcion sin ejecutar lo demas
	}
	// Para grabar la tabla en un archivo.
	fmt.Fprintln(archivo, texto)
	// Cada vez que se abra un archivo se tiene que cerrar
	archivo.Close()

}

// Concatena las tablas
func Suma_tabla() {
	var texto string
	texto = Ejercicios.Ingresar_numero()
	// if func_Append(filename, texto) == false {
	// Aplicando buenas practicas de programacion.
	if !func_Append(filename, texto) {
		fmt.Println("Error al agregar informacion al archivo ... ")
	}

}

func func_Append(filename string, texto string) (estado bool) {
	// "|" = Es un separador de las instrucciones solamente.
	archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al escribir en el archivo ... " + err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)

	if err != nil {
		fmt.Println("Error al escribir en el WriteString ... " + err.Error())
		return false
	}

	archivo.Close()

	return true
}

func Leo_archivo() {
	archivo, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error a la lectura del archivo !!!", err.Error())
	}

	// Lo convertir de Bytyes a cadena.
	fmt.Println(string(archivo))
}

func Leo_archivo2() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error a la lectura del archivo !!!", err.Error())
	}

	// Para leer el archivo.
	textos := bufio.NewScanner(archivo)
	for textos.Scan() {
		registro := textos.Text()
		fmt.Println("> ", registro)
	}

}

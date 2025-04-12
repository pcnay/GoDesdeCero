package Variables

import (
	"runtime"
	"fmt"
)

func Condicionales_If(){
	os := runtime.GOOS // Obtiene el Sistema Operativo
	
	if os =="linux"{
		fmt.Println("El sistema operativo es Linux")
	}
	if os:= runtime.GOOS; os == "darwin"{
		fmt.Println("El sistema operativo es Linux")		
	}
}

package Variables

import (
	"runtime"
	"fmt"
)

func Switch(){
	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Es linux")
		case "darwin":
			fmt.Println("Es Darwin")
		default:
			fmt.Printf("%s \n", os)
	}

}

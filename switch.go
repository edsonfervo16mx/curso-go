package main

import "fmt"
import "time"

func main(){

//momento := time.Now()
hoy:= time.Now().Weekday()

	switch hoy {
		case 0:
			fmt.Println("Hoy es domingo")
		case 1:
			fmt.Println("Hoy es lunes")
		case 2:
			fmt.Println("Hoy es martes")
		case 3:
			fmt.Println("Hoy es miercoles")
		case 4:
			fmt.Println("Hoy es jueves")
		case 5:
			fmt.Println("Hoy es viernes")
		case 6:
			fmt.Println("Hoy es sabado")
	}

}
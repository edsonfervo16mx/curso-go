package main

import "fmt"

func main(){
	//primera forma de declarar y definir valores a un arreglo
	var peliculas [3] string

	peliculas[0] = "Harry potter 1"
	peliculas[1] = "El gran gatsby"
	peliculas[2] = "Rapido y Furioso Reto Tokio"

	fmt.Println(peliculas)

	fmt.Println(peliculas[1])

	//Otra forma de declarar un arreglo

	libros := [3]string{"La marca de los 4","El sabueso de los barkskerville","Romeo y Julieta"}

	fmt.Println(libros)


	var ciudades [3][2]string

	ciudades[0][0] = "Comalcalco"
	ciudades[0][1] = "Centro"
	ciudades[1][0] = "Huimanguillo"
	ciudades[1][1] = "Jalpa"
	ciudades[2][0] = "Teapa"
	ciudades[2][1] = "Balancan"

	fmt.Println(ciudades)

	fmt.Println(ciudades[0])

	fmt.Println(ciudades[0][1])

}

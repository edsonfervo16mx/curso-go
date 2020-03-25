package main

import "fmt"
import "io/ioutil"
import "os"

func main(){
	fmt.Println("Lector")

	//nuevo_texto := []byte(os.Args [1])
	//fmt.Println(string(nuevo_texto))
	nuevo_texto := os.Args [1]+"\n"

	//Me sustituye todo el texto del archivo por el parametro
	//escribir := ioutil.WriteFile("fichero.txt",nuevo_texto,0777)


	//abre el fichero y agrega el parametro al contenido
	fichero, err := os.OpenFile("fichero.txt",os.O_APPEND,0777)
	showError(err)
	escribir, err := fichero.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError(err)
	fichero.Close()

	//showError(escribir)


	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")

	showError(errorDeFichero)

	fmt.Println(string(texto))
	//fmt.Println("****************")
	//fmt.Println(string(texto[3]))
}

func showError(e error){
	if(e != nil){
		panic(e)
	}
}
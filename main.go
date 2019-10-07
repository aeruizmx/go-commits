package main

import (
	"fmt"
	/*	"./flow"
		"./name"
		"./numbers"
		"./strings2"
		"./structs"*/
	"./maps"
)

const helloWorld string = "Hola %s %s, bievenido al fascinante mundo de Go. \n"
const testConst = "Test"

func main() {
	/*lastname := "Ruiz"
	firstName := name.GetName()
	fmt.Printf(helloWorld, firstName, lastname)

	fmt.Println(firstName)

	var firstNumber float64
	var secondNumber float64

	fmt.Print("Ingresa un numero: ")
	fmt.Scanf("%f", &firstNumber)
	fmt.Print("Ingresa otro numero: ")
	fmt.Scanf("%f", &secondNumber)

	var suma = numbers.Suma(firstNumber, secondNumber)
	fmt.Printf("El resultado de la suma es %f \n", suma)
	var resta = numbers.Resta(firstNumber, secondNumber)
	fmt.Printf("El resultado de la resta es %f \n", resta)

	a, b, c := numbers.GetVariables()
	fmt.Println(a, b, c)

	f32, f64 := numbers.GetFloat()
	fmt.Println(f32, f64)

	stringUFT8 := name.GetUnicode()
	fmt.Println("cadena con utf8: ", stringUFT8)
	fmt.Println(string("Hola"[2]))
	fmt.Println("Cantidad de letras: ", len("Hola"))

	structs.GetArray()

	structs.GetSlice()

	flow.IfTest()

	adivinarNumero()

	flow.ForTest()

	strings2.Strings2()

	flow.SwitchTest()*/
	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetData("Andres"))
}

func getCompleteName() (string, string) {
	var name string
	var lastname string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Ingresa tu Apellido: ")
	fmt.Scanf("%s", &lastname)
	return name, lastname
}

func adivinarNumero() {
	numeroBuscado := 10
	var numeroIntroducido int
	fmt.Println("Adivina el numero, esta entre el 1 y 10")
	fmt.Scanf("%d", &numeroIntroducido)
	if numeroIntroducido == numeroBuscado {
		fmt.Println("Adivinaste el numero")
	} else {
		fmt.Println("Suerte para la proxima... :(")
	}
}

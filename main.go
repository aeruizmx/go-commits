package main

import "fmt"

const helloWorld string = "Hola %s %s, bievenido al fascinante mundo de Go. \n"
const testConst = "Test"

func main() {
	name, lastname := getCompleteName()
	fmt.Printf(helloWorld, name, lastname)

	var firstNumber int
	var secondNumber int

	fmt.Print("Ingresa un numero: ")
	fmt.Scanf("%d", &firstNumber)
	fmt.Print("Ingresa otro numero: ")
	fmt.Scanf("%d", &secondNumber)

	var suma = suma(firstNumber, secondNumber)
	fmt.Printf("El resultado de la suma es %d \n", suma)
	var resta = resta(firstNumber, secondNumber)
	fmt.Printf("El resultado de la resta es %d \n", resta)

	a, b, c := getVariables()
	fmt.Println(a, b, c)

	f32, f64 := getFloat()
	fmt.Println(f32, f64)
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
func getVariables() (int, int32, int64) {
	return 1, 241700000, 1919191181818181
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}

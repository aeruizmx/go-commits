package main

import "fmt"

const helloWorld string = "Hola %s %s, bievenido al fascinante mundo de Go. \n"
const testConst = "Test"

func main() {
	name, lastname := getCompleteName()
	fmt.Printf(helloWorld, name, lastname)

	var firstNumber float64
	var secondNumber float64

	fmt.Print("Ingresa un numero: ")
	fmt.Scanf("%f", &firstNumber)
	fmt.Print("Ingresa otro numero: ")
	fmt.Scanf("%f", &secondNumber)

	var suma = suma(firstNumber, secondNumber)
	fmt.Printf("El resultado de la suma es %f \n", suma)
	var resta = resta(firstNumber, secondNumber)
	fmt.Printf("El resultado de la resta es %f \n", resta)

	a, b, c := getVariables()
	fmt.Println(a, b, c)

	f32, f64 := getFloat()
	fmt.Println(f32, f64)

	stringUFT8 := getUnicode()
	fmt.Println("cadena con utf8: ", stringUFT8)
	fmt.Println(string("Hola"[2]))
	fmt.Println("Cantidad de letras: ", len("Hola"))
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

func suma(a float64, b float64) float64 {
	return a + b
}

func resta(a float64, b float64) float64 {
	return a - b
}

func getUnicode() string {
	return "がっこうへいく"
}

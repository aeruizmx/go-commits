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

	getArray()

	getSlice()

	ifTest()

	adivinarNumero()

	forTest()
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

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "Array"
	arr1[1] = "Array 2"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "silce", "1")
	fmt.Println(slice1)
}

func ifTest() {
	var number = 0
	fmt.Println("Ingresa un numero del 1 al 10")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entro al condicional")
	}
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

func forTest() {

	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[FOR] Solo con una condicion de c > 0", c)
	}

	s := 1000
	for {
		s -= 1
		if s == 0 {
			fmt.Println("Termina el for infinito")
			break
		}
	}
}

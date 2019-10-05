package main

import "fmt"

const helloWorld string = "Hola %s %s, bievenido al fascinante mundo de Go. \n"
const testConst = "Test"

func main() {
	name := getName()
	lastname := "<Modificar con mi apellido>"
	number := sum(100, 100)
	a, b, c := getVariables()

	fmt.Printf(helloWorld, name, lastname)
	fmt.Println("Hola mundo")
	fmt.Println(number, a, b, c)
}

func getName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}

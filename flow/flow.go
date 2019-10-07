package flow

import "fmt"

//SwitchTest solicita un numero del 1 al 10 y lo evalua
func SwitchTest() {
	var number = 0
	fmt.Println("[SWITCH] Ingresa un numero del 1 al 10")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("El numero es 1")
	default:
		fmt.Println("El numero no es uno")
	}

	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}
}

//IfTest solicita un numero del 1 al 10 y lo evalua
func IfTest() {
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

//ForTest es una funcion que imprime una secuencia ciclada
func ForTest() {

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
		s--
		if s == 0 {
			fmt.Println("Termina el for infinito")
			break
		}
	}
}

package name

import "fmt"

// GetName obtiene y retorna el nombre una persona
func GetName() string {
	var name string
	name = "Sin nombre"
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

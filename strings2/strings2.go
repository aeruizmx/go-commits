package strings2

import (
	"fmt"
	"strings"
)

//Strings2 es una funcion que opera sobre una cadena de texto (Mayusculas, Minusculas y reemplazo)
func Strings2() {
	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, " "))
}

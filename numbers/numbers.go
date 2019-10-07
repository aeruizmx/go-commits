package numbers

//GetVariables devuelve 3 numeros enteros
func GetVariables() (int, int32, int64) {
	return 1, 241700000, 1919191181818181
}

//GetFloat devuelve 2 numeros lotantes
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

//Suma suma 2 numeros enteros y devuelve resultado
func Suma(a float64, b float64) float64 {
	return a + b
}

//Resta resta 2 numeros enteros y devuelve resultado
func Resta(a float64, b float64) float64 {
	return a - b
}

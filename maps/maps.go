package maps

//GetMap devuelve un map con llave tipo string y valores enteros
func GetMap() map[string]int {
	//mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Juan":   18,
		"Maria":  20,
		"Andres": 32,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100

	delete(mapTest, "llave2")
	delete(mapTest, "llave1")

	return mapTest
}

//GetData devuelve nombre y edad del mapa
func GetData(name string) int {
	//mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Juan":   18,
		"Maria":  20,
		"Andres": 32,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100

	delete(mapTest, "llave2")
	delete(mapTest, "llave1")

	return mapTest[name]
}

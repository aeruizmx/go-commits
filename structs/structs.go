package structs

import "fmt"

//GetArray imprime dos arrays estaticos
func GetArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "Array"
	arr1[1] = "Array 2"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

//GetSlice imprime un array dinamico
func GetSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "silce", "1")
	fmt.Println(slice1)
}

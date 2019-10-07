package structs

import "fmt"

//Struct PlatziCourse
type PlatziCourse struct {
	Name    string
	Slug    string
	Teacher PlatziInstructor
	Skills  []string
}

//Struct PlatziInstructor
type PlatziInstructor struct {
	Name           string
	ProfilePicture string
	TeachesCourses []string
}

//Subscribe es un método de PlatziCareer
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s\n", name, p.Name)
}

//Struct PlatziCareer es una estructura
type PlatziCareer struct {
	Career string
	Course PlatziCourse
}

//Subscribe es un método de PlatziCareer
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado a la carrera %s\n", name, p.Career)
}

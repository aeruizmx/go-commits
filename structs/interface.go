package structs

//Paltzi es una interface de PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func InterfaceTest() {
	platziCareer := PlatziCareer{
		Career: "Desarrollador GO",
		Course: PlatziCourse{Name: "Go", Slug: "/cursos/go", Teacher: PlatziInstructor{Name: "Yohan Graterol", ProfilePicture: "yohan-graterol.jpg", TeachesCourses: []string{"Intruducción a Go", "Go", "Bases de Datos con MongoDB"}}, Skills: []string{"Desarrollo Backend", "Desarrollo Templates"}},
	}
	platziCourse := PlatziCourse{Name: "Go", Slug: "/cursos/go", Teacher: PlatziInstructor{Name: "Yohan Graterol", ProfilePicture: "yohan-graterol.jpg", TeachesCourses: []string{"Intruducción a Go", "Go", "Bases de Datos con MongoDB"}}, Skills: []string{"Desarrollo Backend", "Desarrollo Templates"}}
	//platziCareer.Subscribe("JoseL")
	callSuscribe(platziCareer)
	callSuscribe(platziCourse)
	//fmt.Println(platziCareer)
}

func callSuscribe(p Platzi) {
	p.Subscribe("ARUIZ")
}

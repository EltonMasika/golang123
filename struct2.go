package main
import ("fmt")
func main () {
	aDoctor := struct{name string}{name:"Jon Pertwee"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Elton Masika"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
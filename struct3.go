package main
import ("fmt")

type Animal struct {
	Name string
	Origin string
}
type Bird struct {
	Animal 
	SpeedKPH float32
	Canfly bool
}
 func main () {
	b := Bird {}
	b.Name ="Chicken"
	b.Origin = "Australia"
	b.SpeedKPH = 36
	b.Canfly  = false
	fmt.Println(b)
	fmt.Println(b.Name)
 }
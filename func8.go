package main
import (
	"fmt"

)
func main (){
	g := greeter {
		greeting: "Helloooo",
		name: "Goooooo",
	}
	g.greet()
}
type greeter struct {
	greeting string
	name string
}	
func (g greter ) greet(){

	fmt.Println(g.greeting, g.name)
}



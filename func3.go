package main
import (
	"fmt"

)
func main (){
	sayGreeting("Hellooo", "Elton")
}
func sayGreeting (greeting string, name string){
	fmt.Println(greeting , name)
}
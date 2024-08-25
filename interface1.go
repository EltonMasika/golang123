package main 
import (
	"fmt"
)
func main {
	var w Writer = ConsoleWriter{}
	w.Write([byte]("Hello Playground"))
}
type writer interface {
	Write ([]byte) (int,error)

}
type ConsoleWriter struct {}
func (cw ConsoleWriter) Write (data []byte) (int,error) {
	n,error := fmt.Println(string(data))
	return n ,err
}
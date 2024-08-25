package main
import (
	"fmt"
)
func main () {
	i := 100
	switch{
    case i<= 10:
	fmt.Println("Less than or equal to ten")
	case i <= 20:
	fmt.Println("Less than or equal to twenty")
default:
	fmt.Println("Greater than Twenty")
}
}

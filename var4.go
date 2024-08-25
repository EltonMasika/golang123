package main
import (

	"fmt"
	"strconv"
)

func main() {
    var j int = 42
	fmt.Printf("%v,%T\n", j ,j)
    var i string
	i = strconv.Itoa(j)
    fmt.Printf("%v,%T\n", i ,i )
}
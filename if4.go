package main
import (
	"fmt"
	"math"
)
func main (){
	myNum := 0.1234567
	if myNum == math.Pow(math.Sqrt(myNum),2) {
		fmt.Println("These are the same")
	}else{
		fmt.Println("These are different")
	}
}

package main
import (
	"fmt"
)
func main (){
	number := 50
	guess := 510
	//if guess <=1 && guess <=100{}
	if guess < 1 || guess >100 {

		fmt.Println("The guess must be between 1 and 100!!!!")
	}
	      
	if guess < number {
		fmt.Println("Too low")

	
	}
	if guess > number {
		fmt.Println ("Too High")
	}
	if guess == number {
		fmt.Println ("You got it")
	}
}
package main
import ("fmt")
func main (){
	a := [] int{1,2,3,4,5,6,7,8,9,10}
	b:= a[:] //slice of all elements
	c := a[3:] //slice of 4 to end
	d := a[:6] //slice of first 6 elements
	e := a[3:6] //slice the 4th,5th and 6th elments
	a[5] =657
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
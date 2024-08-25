package main
import ("fmt")
func main (){
	a := [] int {}
	fmt.Println(a)
	fmt.Println("Length: ",len(a))
	fmt.Println("Capacity: ",cap(a))
	a = append(a,2,3,4,5)
	fmt.Println(a)
	fmt.Println("Length: ",len(a))
	fmt.Println("Capacity: ",cap(a))
}
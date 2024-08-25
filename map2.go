package main
import ("fmt")
func main() {
	statepopulations :=make(map[string]int)
	statepopulations = map[string]int{
		"Nairobi" : 35686904,
		"Kisumu" : 10545487,
		"Nakuru" : 486740273,
		"Mombasa" : 78483639,
		"Eldoret" : 73450494,
	}
	fmt.Println(statepopulations)

}
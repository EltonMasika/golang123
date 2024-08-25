package main

import (

	"fmt"
)
type Doctor struct {
	number int
	actorName string 
	companions [] string
}

func main () {
	aDoctor := Doctor{
		number: 3,
		actorName : "Jon PertWee",
		companions : [] string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jessica Parker",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
}
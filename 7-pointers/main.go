package main

import (
	"fmt"
)

func main() {
	var i string = "Nashe"
	//save the addres of i in j
	j := &i
	// copy the value of i in a new addres K
	k := i
	fmt.Println(j)
	// read the addres that contains J
	fmt.Println(*j)
	i = "Nashe updated"
	fmt.Println(j)
	fmt.Println(*j)
	fmt.Println(k)

}

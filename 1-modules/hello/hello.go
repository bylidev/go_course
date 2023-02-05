package main

import (
	"fmt"

	"byli.dev/greetings"
)

func main() {
	message := greetings.Hello("byli!")
	fmt.Println(message)
}

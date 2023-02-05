package main

import (
	"fmt"
	"runtime"
)

func main() {
	goForLoop()
	goWhile()
	goShortIf()
	goSwitch()
	goDefer()
}

func goForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func goWhile() {
	sum := 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)
}

func goShortIf() {
	if myVar := 1; myVar == 1 {
		fmt.Println(myVar)
	}
}

func goSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

/*
Stacking defers
Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed in last-in-first-out order.
*/
func goDefer() {
	fmt.Println("first message")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("second message")
}

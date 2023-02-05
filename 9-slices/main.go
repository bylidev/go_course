package main

import (
	"fmt"
	"strings"
)

func main() {
	var j [10]int
	j[0] = 1
	j[1] = 2
	fmt.Println(j)

	//inline declaration
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slice A slice does not store any data, it just describes a section of an underlying array.
	//WARNING MUTABLE REFERENCE TO ORIGINAL ARRAY
	// in math is like (A,B]

	slice := primes[1:4]
	fmt.Println(slice)
	fmt.Println(primes[:2])
	fmt.Println(primes[2:])
	fmt.Println(primes[5:])

	//make function fixed size with zero values
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	matrix()
	appendIsListLike()
	capVsLenFunction()
	rangeFunction()
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func matrix() {
	// Create a tic-tac-toe board.
	fmt.Println("MATRIX")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendIsListLike() {
	var myArr []string
	fmt.Println(append(myArr, "listLike"))
	// We can add more than one element at a time.
	fmt.Println(append(myArr, "other", "etc"))
	//not mutable function
	fmt.Println(myArr)
	myArr = append(myArr, "other", "etc")
	fmt.Println(myArr)

}

func capVsLenFunction() {
	//cap is the capacity of the array, the append function
	// duplicate initial array size if it needed.
	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}

}

func rangeFunction() {
	//iterate over an array
	fmt.Println("Iterate array with range")

	var myArr = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range myArr {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//alternative
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

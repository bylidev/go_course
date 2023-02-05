package main

import (
	"fmt"
)

type MyExportStruct struct {
	X int
	Y int
}

func main() {
	a := MyExportStruct{1, 2}
	fmt.Println(a.X)
	b := &a
	//this expresion is redundant
	(*b).X = 3
	fmt.Println(a.X)
	//its the same as
	b.X = 4
	fmt.Println(a.X)
	// pointer to struct
	p := &MyExportStruct{1, 2} // has type *Vertex
	fmt.Println(p)
	fmt.Println(*p)

	//declare inline with values
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}

package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	//literals
	m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	//cleanner
	m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
	a, b := m2["gooRgle"]
	fmt.Println(a, b)
	elem, ok := m2["Google"]
	fmt.Println(elem, ok)
	//delete
	delete(m2, "Google")
	elem2, ok2 := m2["Google"]
	fmt.Println(elem2, ok2)

}

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	for _, v := range strings.Split(s, " ") {
		a[v] = 1
	}
	return a
}

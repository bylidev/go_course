package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for y := range p {
		for x := range p[y] {
			p[y][x] = (uint8(x) + uint8(y))
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}

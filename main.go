package main

import (
	"fmt"
	"os"
)

const (
	width  = 30
	height = 20
)

const (
	right direction = iota
	down
	left
	up
)

var directions = [...]direction{right, down, left, up}

type vector struct {
	X int
	Y int
}

type direction int

type board [width * height]int

func (d direction) get_vector() vector {
	switch d {
	case up:
		return vector{0, 1}
	case down:
		return vector{0, -1}
	case left:
		return vector{-1, 0}
	case right:
		return vector{1, 0}
	}

	panic("Tried to get vector on non existent direction")
}

func (d direction) print_command() {
	switch d {
	case up:
		fmt.Println("UP")
	case down:
		fmt.Println("DOWN")
	case left:
		fmt.Println("LEFT")
	case right:
		fmt.Println("RIGHT")
	}
}

func (v vector) get_index() int {
	return v.Y*width + v.X
}

func (v vector) add(x vector) vector {
	return vector{v.X + x.X, v.Y + v.X}
}

func (v *vector) add_self(x vector) {
	v.X += x.X
	v.Y += x.Y
}

func (b *board) set(pos vector, value int) {
	b[pos.get_index()] = value
}

func (b board) get(pos vector) int {
	index := pos.get_index()

	if index > len(b) {
		return 5
	} else {
		return b[index]
	}
}

func (b board) is_safe(pos vector) bool {
	return b.get(pos) == 0
}

func (b *board) update(input []vector) {
	for i := 0; i < len(input); i++ {
		b.set(input[i], i+1)
	}
}

func (b board) debug_print() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := b[vector{x, y}.get_index()]
			fmt.Fprintf(os.Stderr, "%d ", value)
		}

		fmt.Fprintln(os.Stderr)
	}
}

var game = board{}
var current_direction = left

func main() {
	for {
		input := get_input()
		game.update(input)

		game.debug_print()

		move()
	}
}

func get_input() []vector {
	var N, P int
	fmt.Scan(&N, &P)

	input := []vector{}

	for i := 0; i < N; i++ {
		var X0, Y0, X1, Y1 int
		fmt.Scan(&X0, &Y0, &X1, &Y1)

		input = append(input, vector{X1, Y1})
	}

	return input
}

func move() {

}

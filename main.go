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
		return vector{0, -1}
	case down:
		return vector{0, 1}
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
	return vector{v.X + x.X, v.Y + x.Y}
}

func (v vector) add_direction(x direction) vector {
	return v.add(x.get_vector())
}

func (v vector) multiply_scalar(x int) vector {
	return vector{v.X * x, v.Y * x}
}

func (b *board) set(pos vector, value int) {
	b[pos.get_index()] = value
}

func (b board) get(pos vector) int {
	if pos.X < 0 || pos.X >= width || pos.Y < 0 || pos.Y >= height {
		return 5
	}

	index := pos.get_index()

	return b[index]
}

func (b board) is_safe(pos vector) bool {
	safe := b.get(pos) == 0

	return safe
}

func (b *board) update(input []vector) {
	for i, vec := range input {
		b.set(vec, i+1)
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

var (
	game = board{}

	current_pos       vector
	current_direction = left
)

func main() {
	for {
		input := get_input()
		game.update(input)

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

		// otherwise we might miss the end of opponent
		input = append(input, vector{X0, Y0})

		input = append(input, vector{X1, Y1})

		if i == P {
			current_pos = vector{X1, Y1}
		}
	}

	return input
}

func move() {
	if game.is_safe(current_pos.add_direction(current_direction)) {
		current_direction.print_command()
		return
	}

	var best_direction direction
	var best_open_count int

	for _, dir := range directions {
		if game.is_safe(current_pos.add_direction(dir)) {
			open_count := open_square_count(current_pos, dir)

			fmt.Fprintf(os.Stderr, "Tested %d, %d open squares\n", dir.get_vector(), open_count)

			if open_count > best_open_count {
				best_direction = dir
				best_open_count = open_count
			}
		}
	}

	current_direction = best_direction

	current_direction.print_command()
}

func open_square_count(pos vector, dir direction) int {
	for count := 0; ; count++ {
		if !game.is_safe(pos.add(dir.get_vector().multiply_scalar(count + 1))) {
			return count
		}
	}
}

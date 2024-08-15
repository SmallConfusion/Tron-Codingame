package main

import (
	"fmt"
	"os"
	"time"
)

const (
	WIDTH  = 30
	HEIGHT = 20
)

// ---------- direction ----------

type Direction int

const (
	RIGHT Direction = iota
	DOWN
	LEFT
	UP
)

var directions = [...]Direction{RIGHT, DOWN, LEFT, UP}

func (d Direction) GetVector() Vector {
	switch d {
	case UP:
		return Vector{0, -1}
	case DOWN:
		return Vector{0, 1}
	case LEFT:
		return Vector{-1, 0}
	case RIGHT:
		return Vector{1, 0}
	}

	panic("Tried to get vector on non existent direction")
}

func (d Direction) print_command() {
	switch d {
	case UP:
		fmt.Println("UP")
	case DOWN:
		fmt.Println("DOWN")
	case LEFT:
		fmt.Println("LEFT")
	case RIGHT:
		fmt.Println("RIGHT")
	}
}

// ---------- vector ----------

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(x Vector) Vector {
	return Vector{v.X + x.X, v.Y + x.Y}
}

func (v Vector) AddDirection(x Direction) Vector {
	return v.Add(x.GetVector())
}

func (v Vector) MultiplyScalar(x int) Vector {
	return Vector{v.X * x, v.Y * x}
}

func (v Vector) GetIndex() int {
	return v.Y*WIDTH + v.X
}

// ---------- board ----------

type Board [WIDTH * HEIGHT]int

func (b *Board) Set(pos Vector, value int) {
	b[pos.GetIndex()] = value
}

func (b Board) Get(pos Vector) int {
	if pos.X < 0 || pos.X >= WIDTH || pos.Y < 0 || pos.Y >= HEIGHT {
		return 5
	}

	index := pos.GetIndex()

	return b[index]
}

func (b Board) IsSafe(pos Vector) bool {
	safe := b.Get(pos) == 0

	return safe
}

func (b *Board) ClearPlayer(player int) {
	for index, value := range b {
		if value == player {
			b[index] = 0
		}
	}
}

func (b Board) DebugPrint() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			value := b[Vector{x, y}.GetIndex()]
			fmt.Fprintf(os.Stderr, "%d ", value)
		}

		fmt.Fprintln(os.Stderr)
	}
}

// ---------- main ----------

var (
	game = Board{}

	currentPos       Vector
	currentDirection = LEFT
)

func main() {
	for {
		start := time.Now()

		handle_input()
		move()

		elapsed := time.Since(start)
		extra := (time.Millisecond * 100) - elapsed

		fmt.Fprintf(os.Stderr, "Execution time: %s\nExtra time: %s", elapsed, extra)
	}
}

func handle_input() []Vector {
	var N, P int
	fmt.Scan(&N, &P)

	input := []Vector{}

	for i := 0; i < N; i++ {
		var X0, Y0, X1, Y1 int
		fmt.Scan(&X0, &Y0, &X1, &Y1)

		if X0 == -1 {
			game.ClearPlayer(i + 1)
			continue
		}

		game.Set(Vector{X0, Y0}, i+1)
		game.Set(Vector{X1, Y1}, i+1)

		if i == P {
			currentPos = Vector{X1, Y1}
		}
	}

	return input
}

func move() {
	var potDirOpen [4]int

	for i, dir := range directions {
		potDirOpen[i] = openSquareCount(currentPos, dir)
	}

	if potDirOpen[currentDirection] <= 1 {
		var max int
		var maxIndex int

		for index, count := range potDirOpen {
			if max < count {
				max = count
				maxIndex = index
			}
		}

		currentDirection = Direction(maxIndex)
	}

	currentDirection.print_command()
}

func openSquareCount(pos Vector, dir Direction) int {
	for count := 0; ; count++ {
		if !game.IsSafe(pos.Add(dir.GetVector().MultiplyScalar(count + 1))) {
			return count
		}
	}
}

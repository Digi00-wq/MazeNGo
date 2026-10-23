package prim

import (
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

type Edges struct {
	Wall   Point
	Target Point
}

type FieldType int

const (
	FieldEmpty   FieldType = iota // 0
	FieldWall                     // 1
	FieldScanned                  // 2
	FieldStart
	FieldEnd
)

var Directions = []Point{
	{X: 0, Y: -1}, // North
	{X: 0, Y: 1},  // South
	{X: -1, Y: 0}, // West
	{X: 1, Y: 0},  // East
}

func CreateMaze(a chan [][]FieldType, width, height int) [][]FieldType {
	if width < 5 {
		width = 5
	}
	if height < 5 {
		height = 5
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	field := make([][]FieldType, width)
	for x := range width {
		field[x] = make([]FieldType, height)
		for y := range height {
			field[x][y] = FieldWall
		}
	}

	startingCell := Point{X: 1, Y: 1}
	field[startingCell.X][startingCell.Y] = FieldEmpty

	var mazeDirections []Edges

	addToList := func(p Point) {
		for _, d := range Directions {
			target := Point{X: p.X + d.X*2, Y: p.Y + d.Y*2}
			if InBounds(target, width, height) && field[target.X][target.Y] == FieldWall {
				wall := Point{X: p.X + d.X, Y: p.Y + d.Y}
				mazeDirections = append(mazeDirections, Edges{Wall: wall, Target: target})
			}
		}
	}

	addToList(startingCell)

	for len(mazeDirections) > 0 {
		idx := rng.Intn(len(mazeDirections))
		edge := mazeDirections[idx]

		mazeDirections[idx] = mazeDirections[len(mazeDirections)-1]
		mazeDirections = mazeDirections[:len(mazeDirections)-1]

		if field[edge.Target.X][edge.Target.Y] == FieldWall {
			field[edge.Wall.X][edge.Wall.Y] = FieldEmpty
			field[edge.Target.X][edge.Target.Y] = FieldEmpty

			addToList(edge.Target)
		}
		time.Sleep(100)
		a <- field
	}

	a <- field
	return field
}

func InBounds(p Point, width, height int) bool {
	return p.X > 0 && p.X < width-1 && p.Y > 0 && p.Y < height-1
}

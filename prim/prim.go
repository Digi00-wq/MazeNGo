package prim

import (
	"math/rand"
	"time"
)

type point struct {
	x, y int
}

type Edges struct {
	wall   point
	target point
}

type FieldType int

const (
	FieldEmpty FieldType = iota // 0
	FieldWall                   // 1
)

func CreateMaze(size int, a chan [][]FieldType) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	field := make([][]FieldType, size)
	for x := range size {
		field[x] = make([]FieldType, size)
		for y := range size {
			field[x][y] = FieldWall
		}
	}

	startingCell := point{x: 1, y: 1}
	field[startingCell.x][startingCell.y] = FieldEmpty

	directions := []point{
		{x: 0, y: -1}, // North
		{x: 0, y: 1},  // South
		{x: -1, y: 0}, // West
		{x: 1, y: 0},  // East
	}

	var mazeDirections []Edges

	addToList := func(p point) {
		for _, d := range directions {
			target := point{x: p.x + d.x*2, y: p.y + d.y*2}
			if inBounds(target, size) && field[target.x][target.y] == FieldWall {
				wall := point{x: p.x + d.x, y: p.y + d.y}
				mazeDirections = append(mazeDirections, Edges{wall: wall, target: target})
			}
		}
	}

	addToList(startingCell)

	for len(mazeDirections) > 0 {
		idx := rng.Intn(len(mazeDirections))
		edge := mazeDirections[idx]

		mazeDirections[idx] = mazeDirections[len(mazeDirections)-1]
		mazeDirections = mazeDirections[:len(mazeDirections)-1]

		if field[edge.target.x][edge.target.y] == FieldWall {
			field[edge.wall.x][edge.wall.y] = FieldEmpty
			field[edge.target.x][edge.target.y] = FieldEmpty

			addToList(edge.target)
		}
	}

	a <- field
}

func inBounds(p point, size int) bool {
	return p.x > 0 && p.x < size-1 && p.y > 0 && p.y < size-1
}

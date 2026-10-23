package dfs

import (
	"MazeNGo/prim"
	"errors"
	"math/rand"
	"slices"
	"time"
)

type thingy struct {
	ways         []prim.Point
	currentPoint prim.Point
}

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		values: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.values = append(s.values, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.values) == 0 {
		return zero, errors.New("stack is empty")
	}

	index := len(s.values) - 1
	item := s.values[index]
	s.values = s.values[:index]

	return item, nil
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func SolveMaze(field [][]prim.FieldType, size int) {
	if len(field) <= 4 || len(field[0]) <= 4 {
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	startingPoint := prim.Point{X: 1, Y: 1}
	endPoint := prim.Point{X: rng.Intn(len(field)-2) + 1, Y: rng.Intn(len(field)-2) + 1}
	for {
		if field[endPoint.X][endPoint.Y] != prim.FieldEmpty {
			endPoint = prim.Point{X: rng.Intn(len(field)-2) + 1, Y: rng.Intn(len(field)-2) + 1}
		} else {
			break
		}

	}
	visited := make([]prim.Point, 0)

	thing := NewStack[thingy]()

	isVisited := func(p prim.Point) bool {
		return slices.Contains(visited, p)
	}

	isValid := func(p prim.Point) bool {
		return p.X >= 0 && p.X < len(field) && p.Y >= 0 && p.Y < len(field[0])
	}

	addThingy := func(p prim.Point) {
		visited = append(visited, p)
		thingle := thingy{ways: make([]prim.Point, 0), currentPoint: p}

		for _, d := range prim.Directions {
			newPoint := prim.Point{X: p.X + d.X, Y: p.Y + d.Y}
			if isValid(newPoint) && field[newPoint.X][newPoint.Y] != prim.FieldWall && !isVisited(newPoint) {
				thingle.ways = append(thingle.ways, newPoint)
			}
		}
		thing.Push(thingle)
	}

	addThingy(startingPoint)
	field[startingPoint.X][startingPoint.Y] = prim.FieldStart
	field[endPoint.X][endPoint.Y] = prim.FieldEnd

	for thing.Len() > 0 {
		returnValue, err := thing.Pop()
		if err != nil {
			break
		}

		if len(returnValue.ways) == 0 {
			continue
		}

		idx := rng.Intn(len(returnValue.ways))
		newPoint := returnValue.ways[idx]

		returnValue.ways[idx] = returnValue.ways[len(returnValue.ways)-1]
		returnValue.ways = returnValue.ways[:len(returnValue.ways)-1]

		if len(returnValue.ways) > 0 {
			thing.Push(returnValue)
		}

		if isVisited(newPoint) {
			continue
		}

		if newPoint == endPoint {
			return
		}
		addThingy(newPoint)

		field[newPoint.X][newPoint.Y] = prim.FieldScanned
	}
}

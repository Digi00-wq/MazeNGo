package main

import (
	"MazeNGo/dfs"
	"MazeNGo/display"
	"MazeNGo/prim"
	"fmt"
	"strconv"
	"time"
)

func main() {
	a := make(chan [][]prim.FieldType)
	b := make(chan int)
	go tutorMain(a)
	display.InitWindow(a, b)
}

func tutorMain(a chan [][]prim.FieldType) {
	time.Sleep(100)

	var value int = 21
	for {
		var input string
		fmt.Scanln(&input)
		if input != "" {
			i, err := strconv.Atoi(input)
			if err != nil {
				continue
			} else {
				value = i
			}
		}
		field := prim.CreateMaze(value, a)
		dfs.SolveMaze(field, len(field))
	}
}

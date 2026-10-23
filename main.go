package main

import (
	"MazeNGo/dfs"
	"MazeNGo/display"
	"MazeNGo/prim"
	"fmt"
	"strconv"
)

func main() {
	field := make(chan [][]prim.FieldType)
	width := make(chan int, 1)
	height := make(chan int, 1)
	go tutorMain(field, width, height)
	display.InitWindow(field, width, height)
}

func tutorMain(field chan [][]prim.FieldType, width chan int, height chan int) {
	var newWidth int = 21
	var newHeight int = 21
	for {
		var input string
		fmt.Scanln(&input)
		if input != "" {
			_, err := strconv.Atoi(input)
			if err != nil {
				continue
			}
		}
		for {
			select {
			case newWidth = <-width:
			case newHeight = <-height:
			default:
				goto create
			}
		}
	create:
		maze := prim.CreateMaze(field, newWidth, newHeight)
		go dfs.SolveMaze(maze)
	}
}

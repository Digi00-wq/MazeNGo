package main

import (
	"MazeNGo/dfs"
	"MazeNGo/display"
	"MazeNGo/prim"
	"fmt"
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
	for {
		var i int
		fmt.Scan(&i)
		field := prim.CreateMaze(i, a)
		dfs.SolveMaze(field, len(field))
	}
}

package main

import (
	"MazeNGo/display"
	"MazeNGo/prim"
	"fmt"
	"time"
)

func main() {
	a := make(chan [][]int)
	b := make(chan int)
	go tutorMain(a)
	display.InitWindow(a, b)
}

func tutorMain(a chan [][]int) {
	time.Sleep(100)
	for {
		var i int
		fmt.Scan(&i)
		prim.CreateMaze(i, a)
	}
}

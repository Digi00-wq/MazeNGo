package main

import (
	"MazeNGo/display"
	"MazeNGo/prim"
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
	prim.CreateMaze(5, a)
}

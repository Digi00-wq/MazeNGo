package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FieldType int

const (
	FieldEmpty FieldType = iota // 0
	FieldWall                   // 1
)

var (
	fontSize int32 = 20
	text           = "Hello world"
	Maze           = [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
)

func InitWindow(a chan [][]int, b chan int) {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(200, 100, "MazeNGo")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		updateField(Maze)
		select {
		case updatedMaze := <-a:
			Maze = updatedMaze
		default:
			updateField(Maze)
		}
	}

}

func getMiddleHeight() int32 {
	return int32(rl.GetRenderHeight()/2) - fontSize/2
}

func getMiddleWidth() int32 {
	return int32(rl.GetRenderWidth()/2) - fontSize*3
}

func updateField(maze [][]int) {

	var basePosX int32 = 300
	var basePosY int32 = 200
	var size int32 = 20
	offset := 10
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	for i := range maze {
		for j := range maze[0] {
			if maze[j][i] == 1 {
				rl.DrawRectangle(basePosX+size*int32(i)+int32(offset*i), basePosY+size*int32(j)+int32(offset*j), size, size, rl.Black)
			} else {

				rl.DrawRectangle(basePosX+size*int32(i)+int32(offset*i), basePosY+size*int32(j)+int32(offset*j), size, size, rl.Green)
			}
		}
	}
	rl.EndDrawing()
}

func SetMaze(nweMaze [][]int) {
	Maze = nweMaze
}

package display

import (
	"MazeNGo/prim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	fontSize int32 = 20
	text           = "Hello world"
	Maze           = [][]prim.FieldType{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
)

func InitWindow(a chan [][]prim.FieldType, b chan int) {
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

func updateField(maze [][]prim.FieldType) {

	var basePosX int32 = 300
	var basePosY int32 = 200
	var size int32 = 20
	offset := 10
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	for i := range maze {
		for j := range maze[0] {
			switch maze[i][j] {
			case prim.FieldWall:
				rl.DrawRectangle(basePosX+size*int32(i)+int32(offset*i), basePosY+size*int32(j)+int32(offset*j), size, size, rl.Black)
			case prim.FieldEmpty:
				rl.DrawRectangle(basePosX+size*int32(i)+int32(offset*i), basePosY+size*int32(j)+int32(offset*j), size, size, rl.Green)
			}
		}
	}
	rl.EndDrawing()
}

func SetMaze(nweMaze [][]prim.FieldType) {
	Maze = nweMaze
}

package display

import (
	"MazeNGo/prim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	fontSize int32 = 20
	text           = "Hello world"
	Maze           = [][]prim.FieldType{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
	toggle         = 0
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

	var basePosX int32 = 50
	var basePosY int32 = 50
	var size int32 = 20
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	for i := range maze {
		for j := range maze[0] {
			switch maze[i][j] {
			case prim.FieldWall:
				rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Black)
			case prim.FieldEmpty:
				rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Red)
			case prim.FieldScanned:
				rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Green)
			case prim.FieldStart:
				rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Gold)
			case prim.FieldEnd:
				thingy := 30
				if toggle <= thingy {
					rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Gold)
					toggle++
				} else {
					rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.White)
					toggle++
				}

				if toggle >= thingy*2 {
					toggle = 0
				}

			}
		}
	}
	rl.EndDrawing()
}

func SetMaze(nweMaze [][]prim.FieldType) {
	Maze = nweMaze
}

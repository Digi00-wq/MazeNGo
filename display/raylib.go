package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	fontSize int32 = 20
	text           = "Hello world"
	maze           = [5][5]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
)

func InitWindow() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(200, 100, "MazeNGo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var basePosX int32 = 300
	var basePosY int32 = 200
	var size int32 = 20

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		for i := range maze {
			for j := range maze[0] {
				if maze[j][i] == 1 {
					rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Black)
				} else {

					rl.DrawRectangle(basePosX+size*int32(i), basePosY+size*int32(j), size, size, rl.Green)
				}
			}
		}
		rl.EndDrawing()

	}

}

func getMiddleHeight() int32 {
	return int32(rl.GetRenderHeight()/2) - fontSize/2
}

func getMiddleWidth() int32 {
	return int32(rl.GetRenderWidth()/2) - fontSize*3
}

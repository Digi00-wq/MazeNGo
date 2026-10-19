package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	fontSize int32 = 20
	text           = "Hello world"
)

func InitWindow() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(200, 100, "MazeNGo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(text, getMiddleWidth(), getMiddleHeight(), fontSize, rl.Black)

		rl.EndDrawing()

	}

}

func getMiddleHeight() int32 {
	return int32(rl.GetRenderHeight()/2) - fontSize/2
}

func getMiddleWidth() int32 {
	return int32(rl.GetRenderWidth()/2) - fontSize*3
}

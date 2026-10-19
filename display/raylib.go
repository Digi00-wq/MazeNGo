package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init_window() {
	rl.InitWindow(1920, 1080, "MazeNGo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Hello world", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()

	}

}

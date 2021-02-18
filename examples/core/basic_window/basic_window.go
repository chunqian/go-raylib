package main

import (
	rl "go-raylib/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

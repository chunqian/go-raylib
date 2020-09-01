package main

import (
	rl "goray/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - mouse input")

	ballPosition := rl.NewVector2(-100, -100)
	ballColor := rl.DarkBlue

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		ballPosition = rl.GetMousePosition()

		if rl.IsMouseButtonPressed(int32(rl.MOUSE_LEFT_BUTTON)) {
			ballColor = rl.Maroon
		} else if rl.IsMouseButtonPressed(int32(rl.MOUSE_MIDDLE_BUTTON)) {
			ballColor = rl.Lime
		} else if rl.IsMouseButtonPressed(int32(rl.MOUSE_RIGHT_BUTTON)) {
			ballColor = rl.DarkBlue
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawCircleV(ballPosition, 40, ballColor)

		rl.DrawText("move ball with mouse and click mouse button to change color", 10, 10, 20, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

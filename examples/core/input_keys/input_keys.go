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

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - keyboard input")

	ballPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			ballPosition.This.X += 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			ballPosition.This.X -= 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_UP)) {
			ballPosition.This.Y -= 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_DOWN)) {
			ballPosition.This.Y += 2.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("move the ball with arrow keys", 10, 10, 20, rl.DarkGray)

		rl.DrawCircleV(ballPosition, 50, rl.Maroon)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

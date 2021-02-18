package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - keyboard input")
	defer rl.CloseWindow()

	ballPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			ballPosition.X += 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			ballPosition.X -= 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_UP)) {
			ballPosition.Y -= 2.0
		}
		if rl.IsKeyDown(int32(rl.KEY_DOWN)) {
			ballPosition.Y += 2.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("move the ball with arrow keys", 10, 10, 20, rl.DarkGray)

		rl.DrawCircleV(ballPosition, 50, rl.Maroon)

		rl.EndDrawing()
	}
}

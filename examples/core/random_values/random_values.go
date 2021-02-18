package main

import (
	rl "go-raylib/raylib"

	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - generate random values")
	defer rl.CloseWindow()

	framesCounter := 0

	randValue := rl.GetRandomValue(-8, 5)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		framesCounter++

		if (framesCounter/120)%2 == 1 {

			randValue = rl.GetRandomValue(-8, 5)
			framesCounter = 0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Every 2 seconds a new random value is generated:", 130, 100, 20, rl.Maroon)

		rl.DrawText(fmt.Sprintf("%d", randValue), 360, 180, 80, rl.LightGray)

		rl.EndDrawing()
	}
}

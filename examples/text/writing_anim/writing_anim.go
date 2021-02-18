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

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - text writing anim")
	defer rl.CloseWindow()

	message := "This sample illustrates a text writing\nanimation effect! Check it out! ;)"

	framesCounter := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_SPACE)) {
			framesCounter += 8
		} else {
			framesCounter++
		}

		if rl.IsKeyPressed(int32(rl.KEY_ENTER)) {
			framesCounter = 0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(rl.TextSubtext(message, 0, int32(framesCounter/10)), 210, 160, 20, rl.Maroon)

		rl.DrawText("PRESS [ENTER] to RESTART!", 240, 260, 20, rl.LightGray)
		rl.DrawText("PRESS [SPACE] to SPEED UP!", 239, 300, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

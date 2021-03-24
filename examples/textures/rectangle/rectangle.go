package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

var (
	MAX_FRAME_SPEED = 15
	MIN_FRAME_SPEED = 1
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [texture] example - texture rectangle")
	defer rl.CloseWindow()

	scarfy := rl.LoadTexture("../textures/resources/scarfy.png")
	defer rl.UnloadTexture(scarfy)

	position := rl.NewVector2(350.0, 280.0)
	frameRec := rl.NewRectangle(0, 0, float32(scarfy.Width)/6, float32(scarfy.Height))
	currentFrame := int32(0)

	framesCounter := int32(0)
	framesSpeed := int32(8)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		framesCounter++

		if framesCounter >= (60 / framesSpeed) {
			framesCounter = 0
			currentFrame++

			if currentFrame > 5 {
				currentFrame = 0
			}
			frameRec.X = float32(currentFrame*scarfy.Width) / 6
		}

		if rl.IsKeyPressed(int32(rl.KEY_RIGHT)) {
			framesSpeed++
		} else if rl.IsKeyPressed(int32(rl.KEY_LEFT)) {
			framesSpeed--
		}

		if framesSpeed > int32(MAX_FRAME_SPEED) {
			framesSpeed = int32(MAX_FRAME_SPEED)
		} else if framesSpeed < int32(MIN_FRAME_SPEED) {
			framesSpeed = int32(MIN_FRAME_SPEED)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(scarfy, 15, 40, rl.White)

		rl.DrawRectangleLines(15, 40, scarfy.Width, scarfy.Height, rl.Lime)
		rl.DrawRectangleLines(15+int32(frameRec.X), 40+int32(frameRec.Y), int32(frameRec.Width), int32(frameRec.Height), rl.Red)

		rl.DrawText("FRAME SPEED: ", 165, 210, 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("%02d FPS", framesSpeed), 575, 210, 10, rl.DarkGray)
		rl.DrawText("PRESS RIGHT/LEFT KEYS to CHANGE SPEED!", 290, 240, 10, rl.DarkGray)

		for i := 0; i < MAX_FRAME_SPEED; i++ {
			if int32(i) < framesSpeed {
				rl.DrawRectangle(int32(250+21*i), 205, 20, 20, rl.Red)
			}
			rl.DrawRectangleLines(int32(250+21*i), 205, 20, 20, rl.Maroon)
		}

		rl.DrawTextureRec(scarfy, frameRec, position, rl.White)

		rl.DrawText("(c) Scarfy sprite by Eiden Marsal", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.EndDrawing()
	}
}

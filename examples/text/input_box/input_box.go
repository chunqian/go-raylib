package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

const MAX_INPUT_CHARS = 9

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - input box")

	name := rl.NewBytes("", MAX_INPUT_CHARS+1)
	letterCount := 0

	textBox := rl.NewRectangle(float32(screenWidth/2-100), 180, 225, 50)
	mouseOnText := false

	framesCounter := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
			mouseOnText = true
		} else {
			mouseOnText = false
		}

		if mouseOnText {

			key := rl.GetKeyPressed()
			for key > 0 {
				if key >= 32 && key <= 125 && letterCount < MAX_INPUT_CHARS {
					name[letterCount] = byte(key)
					letterCount++
				}

				key = rl.GetKeyPressed()
			}

			if rl.IsKeyPressed(int32(rl.KEY_BACKSPACE)) {
				letterCount--
				if letterCount < 0 {
					letterCount = 0
				}
				name[letterCount] = byte(0)
			}
		}

		if mouseOnText {
			framesCounter++
		} else {
			framesCounter = 0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("PLACE MOUSE OVER INPUT BOX!", 240, 140, 20, rl.Gray)

		rl.DrawRectangleRec(textBox, rl.LightGray)
		if mouseOnText {
			rl.DrawRectangleLines(
				int32(textBox.X),
				int32(textBox.Y),
				int32(textBox.Width),
				int32(textBox.Height),
				rl.Red,
			)
		} else {
			rl.DrawRectangleLines(
				int32(textBox.X),
				int32(textBox.Y),
				int32(textBox.Width),
				int32(textBox.Height),
				rl.DarkGray,
			)
		}

		rl.DrawText(string(name), int32(textBox.X+5), int32(textBox.Y+8), 40, rl.Maroon)

		rl.DrawText(fmt.Sprintf("INPUT CHARS: %d/%d", letterCount, MAX_INPUT_CHARS), 315, 250, 20, rl.DarkGray)

		if mouseOnText {
			if letterCount < MAX_INPUT_CHARS {
				if ((framesCounter / 20) % 2) == 0 {
					rl.DrawText("_",
						int32(textBox.X)+8+rl.MeasureText(string(name), 40),
						int32(textBox.Y+12),
						40,
						rl.Maroon,
					)
				}
			}
		}
		rl.EndDrawing()
	}
}

func IsAnyKeyPressed() bool {
	keyPressed := false
	key := rl.GetKeyPressed()

	if key >= 32 && key <= 126 {
		keyPressed = true
	}
	return keyPressed
}

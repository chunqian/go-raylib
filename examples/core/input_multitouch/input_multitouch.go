package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

var MAX_TOUCH_POINTS = 10

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - input multitouch")
	defer rl.CloseWindow()

	ballPosition := rl.NewVector2(-100, -100)
	ballColor := rl.Beige

	touchCounter := 0
	touchPosition := rl.NewVector2(0, 0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		ballPosition = rl.GetMousePosition()

		ballColor = rl.Beige

		if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {
			ballColor = rl.Maroon
		}
		if rl.IsMouseButtonDown(int32(rl.MOUSE_MIDDLE_BUTTON)) {
			ballColor = rl.Lime
		}
		if rl.IsMouseButtonDown(int32(rl.MOUSE_RIGHT_BUTTON)) {
			ballColor = rl.DarkBlue
		}

		if rl.IsMouseButtonPressed(int32(rl.MOUSE_LEFT_BUTTON)) {
			touchCounter = 10
		}
		if rl.IsMouseButtonPressed(int32(rl.MOUSE_MIDDLE_BUTTON)) {
			touchCounter = 10
		}
		if rl.IsMouseButtonPressed(int32(rl.MOUSE_RIGHT_BUTTON)) {
			touchCounter = 10
		}

		if touchCounter > 0 {
			touchCounter--
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := 0; i < MAX_TOUCH_POINTS; i++ {

			touchPosition = rl.GetTouchPosition(int32(i))

			if touchPosition.This.X >= 0 && touchPosition.This.Y >= 0 {

				rl.DrawCircleV(touchPosition, 34, rl.Orange)
				rl.DrawText(fmt.Sprintf("%d", i), int32(touchPosition.This.X)-10, int32(touchPosition.This.Y)-70, 40, rl.Black)
			}
		}

		rl.DrawCircleV(ballPosition, float32(30+(touchCounter*3)), ballColor)

		rl.DrawText("move ball with mouse and click mouse button to change color", 10, 10, 20, rl.DarkGray)
		rl.DrawText("touch the screen at multiple locations to get multiple balls", 10, 30, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}

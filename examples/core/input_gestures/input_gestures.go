package main

import (
	rl "goray/raylib"

	"runtime"
)

var MAX_GESTURE_STRINGS = 20

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - input gestures")

	touchPosition := rl.NewVector2(0, 0)
	touchArea := rl.NewRectangle(220, 10, float32(screenWidth-230), float32(screenHeight-20))

	gesturesCount := int32(0)
	var gestureStrings [32]string

	currentGesture := int32(rl.GESTURE_NONE)
	lastGesture := int32(rl.GESTURE_NONE)

	// rl.SetGesturesEnabled(0b0000000000001001);   // Enable only some gestures to be detected

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		lastGesture = currentGesture

		currentGesture = rl.GetGestureDetected()
		touchPosition = rl.GetTouchPosition(0)

		if rl.CheckCollisionPointRec(touchPosition, touchArea) && currentGesture != int32(rl.GESTURE_NONE) {
			if currentGesture != lastGesture {
				switch rl.GestureType(currentGesture) {
				case rl.GESTURE_TAP:
					gestureStrings[gesturesCount] = "GESTURE TAP"
					break
				case rl.GESTURE_DOUBLETAP:
					gestureStrings[gesturesCount] = "GESTURE DOUBLETAP"
					break
				case rl.GESTURE_HOLD:
					gestureStrings[gesturesCount] = "GESTURE HOLD"
					break
				case rl.GESTURE_DRAG:
					gestureStrings[gesturesCount] = "GESTURE DRAG"
					break
				case rl.GESTURE_SWIPE_RIGHT:
					gestureStrings[gesturesCount] = "GESTURE SWIPE RIGHT"
					break
				case rl.GESTURE_SWIPE_LEFT:
					gestureStrings[gesturesCount] = "GESTURE SWIPE LEFT"
					break
				case rl.GESTURE_SWIPE_UP:
					gestureStrings[gesturesCount] = "GESTURE SWIPE UP"
					break
				case rl.GESTURE_SWIPE_DOWN:
					gestureStrings[gesturesCount] = "GESTURE SWIPE DOWN"
					break
				case rl.GESTURE_PINCH_IN:
					gestureStrings[gesturesCount] = "GESTURE PINCH IN"
					break
				case rl.GESTURE_PINCH_OUT:
					gestureStrings[gesturesCount] = "GESTURE PINCH OUT"
					break
				default:
					break
				}

				gesturesCount++

				if gesturesCount >= int32(MAX_GESTURE_STRINGS) {
					for i := 0; i < MAX_GESTURE_STRINGS; i++ {
						gestureStrings[i] = ""
					}
					gesturesCount = 0
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleRec(touchArea, rl.Gray)
		rl.DrawRectangle(225, 15, screenWidth-240, screenHeight-30, rl.RayWhite)

		rl.DrawText("GESTURES TEST AREA", screenWidth-270, screenHeight-40, 20, rl.Fade(rl.Gray, 0.5))

		for i := int32(0); i < gesturesCount; i++ {

			if i%2 == 0 {
				rl.DrawRectangle(10, 30+20*i, 200, 20, rl.Fade(rl.LightGray, 0.5))
			} else {
				rl.DrawRectangle(10, 30+20*i, 200, 20, rl.Fade(rl.LightGray, 0.3))
			}

			if i < gesturesCount-1 {
				rl.DrawText(gestureStrings[i], 35, 36+20*i, 10, rl.DarkGray)
			} else {
				rl.DrawText(gestureStrings[i], 35, 36+20*i, 10, rl.Maroon)
			}
		}

		rl.DrawRectangleLines(10, 29, 200, screenHeight-50, rl.Gray)
		rl.DrawText("DETECTED GESTURES", 50, 15, 10, rl.Gray)

		if currentGesture != int32(rl.GESTURE_NONE) {
			rl.DrawCircleV(touchPosition, 30, rl.Maroon)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

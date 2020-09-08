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

	rl.InitWindow(screenWidth, screenHeight, "raylib [text] example - sprite font loading")
	defer rl.CloseWindow()

	msg := [3]string{
		"THIS IS A custom SPRITE FONT...",
		"...and this is ANOTHER CUSTOM font...",
		"...and a THIRD one! GREAT! :D",
	}

	font := [3]rl.Font{
		rl.LoadFont("../text/resources/custom_mecha.png"),
		rl.LoadFont("../text/resources/custom_alagard.png"),
		rl.LoadFont("../text/resources/custom_jupiter_crash.png"),
	}
	defer func() {
		for i := range font {
			rl.UnloadFont(font[i])
		}
	}()

	fontPosition := [3]rl.Vector2{
		rl.NewVector2(
			float32(screenWidth)/2-rl.MeasureTextEx(font[0], msg[0], float32(font[0].BaseSize), -3).X/2,
			float32(screenHeight/2-font[0].BaseSize/2-80),
		),
		rl.NewVector2(
			float32(screenWidth)/2-rl.MeasureTextEx(font[1], msg[1], float32(font[1].BaseSize), -2).X/2,
			float32(screenHeight/2-font[1].BaseSize/2-10),
		),
		rl.NewVector2(
			float32(screenWidth)/2-rl.MeasureTextEx(font[2], msg[2], float32(font[2].BaseSize), -3).X/2,
			float32(screenHeight/2-font[2].BaseSize/2+50),
		),
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTextEx(font[0], msg[0], fontPosition[0], float32(font[0].BaseSize), -3, rl.White)
		rl.DrawTextEx(font[1], msg[1], fontPosition[1], float32(font[1].BaseSize), -2, rl.White)
		rl.DrawTextEx(font[2], msg[2], fontPosition[2], float32(font[2].BaseSize), 2, rl.White)

		rl.EndDrawing()
	}
}

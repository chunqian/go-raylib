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

	rl.InitWindow(screenWidth, screenHeight, "raylib [audio] example - sound loading and playing")
	rl.InitAudioDevice()

	fxWav := rl.LoadSound("../audio/resources/sound.wav")
	fxOgg := rl.LoadSound("../audio/resources/target.ogg")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			rl.PlaySound(fxWav)
		}
		if rl.IsKeyPressed(int32(rl.KEY_ENTER)) {
			rl.PlaySound(fxOgg)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Press SPACE to PLAY the WAV sound!", 200, 180, 20, rl.LightGray)
		rl.DrawText("Press ENTER to PLAY the OGG sound!", 200, 220, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.UnloadSound(fxWav)
	rl.UnloadSound(fxOgg)

	rl.CloseAudioDevice()

	rl.CloseWindow()
}

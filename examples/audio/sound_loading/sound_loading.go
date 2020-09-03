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
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	fxWav := rl.LoadSound("../audio/resources/sound.wav")
	defer rl.UnloadSound(fxWav)

	fxOgg := rl.LoadSound("../audio/resources/target.ogg")
	defer rl.UnloadSound(fxOgg)

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
}

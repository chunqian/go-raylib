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

	rl.InitWindow(screenWidth, screenHeight, "raylib [audio] example - music playing (streaming)")
	rl.InitAudioDevice()

	music := rl.LoadMusicStream("../audio/resources/country.mp3")

	rl.PlayMusicStream(music)

	timePlayed := float32(0)
	pause := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateMusicStream(music)

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			rl.StopMusicStream(music)
			rl.PlayMusicStream(music)
		}

		if rl.IsKeyPressed(int32(rl.KEY_P)) {
			pause = !pause

			if pause {
				rl.PauseMusicStream(music)
			} else {
				rl.ResumeMusicStream(music)
			}
		}

		timePlayed = rl.GetMusicTimePlayed(music) / rl.GetMusicTimeLength(music) * 400

		if timePlayed > 400 {
			rl.StopMusicStream(music)
		}

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.DrawText("MUSIC SHOULD BE PLAYING!", 255, 150, 20, *rl.LightGray)

		rl.DrawRectangle(200, 200, 400, 12, *rl.LightGray)
		rl.DrawRectangle(200, 200, int32(timePlayed), 12, *rl.Maroon)
		rl.DrawRectangleLines(200, 200, 400, 12, *rl.Gray)

		rl.DrawText("PRESS SPACE TO RESTART MUSIC", 215, 250, 20, *rl.LightGray)
		rl.DrawText("PRESS P TO PAUSE/RESUME MUSIC", 208, 280, 20, *rl.LightGray)

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(music)

	rl.CloseAudioDevice()

	rl.CloseWindow()
}

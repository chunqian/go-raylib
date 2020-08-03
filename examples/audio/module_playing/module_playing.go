package main

import (
	rl "goray/raylib"
	"runtime"
)

var MAX_CIRCLES = 64

type CircleWave struct {
	position rl.Vector2
	radius   float32
	alpha    float32
	speed    float32
	color    rl.Color
}

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [audio] example - module playing (streaming)")

	rl.InitAudioDevice()

	colors := []*rl.Color{
		rl.Orange,
		rl.Red,
		rl.Gold,
		rl.Lime,
		rl.Blue,
		rl.Violet,
		rl.Brown,
		rl.LightGray,
		rl.Pink, rl.Yellow,
		rl.Green,
		rl.SkyBlue,
		rl.Purple,
		rl.Beige}
	circles := make([]CircleWave, MAX_CIRCLES)

	for i := (MAX_CIRCLES - 1); i >= 0; i-- {
		circles[i].alpha = 0.0
		circles[i].radius = float32(rl.GetRandomValue(10, 40))

		px := float32(rl.GetRandomValue(int32(circles[i].radius), screenWidth-int32(circles[i].radius)))
		py := float32(rl.GetRandomValue(int32(circles[i].radius), screenHeight-int32(circles[i].radius)))
		circles[i].position = rl.Vector2{X: px, Y: py}
		circles[i].speed = float32(rl.GetRandomValue(1, 100)) / 2000.0
		circles[i].color = *colors[rl.GetRandomValue(0, 13)]
	}

	music := rl.LoadMusicStream("../audio/resources/mini1111.xm")
	// if GO struct mapping with C struct please use Convert() to assign value
	// music.Looping = true
	music.Convert().Looping = true

	rl.PlayMusicStream(music)

	timePlayed := float32(0.0)
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

		timePlayed = rl.GetMusicTimePlayed(music) / rl.GetMusicTimeLength(music) * float32(screenWidth-40)

		for i := (MAX_CIRCLES - 1); (i >= 0) && !pause; i-- {
			circles[i].alpha += circles[i].speed
			circles[i].radius += circles[i].speed * 10.0

			if circles[i].alpha > 1.0 {
				circles[i].speed *= -1
			}
			if circles[i].alpha <= 0.0 {
				circles[i].alpha = 0.0
				circles[i].radius = float32(rl.GetRandomValue(10, 40))
				px := float32(rl.GetRandomValue(int32(circles[i].radius), screenWidth-int32(circles[i].radius)))
				py := float32(rl.GetRandomValue(int32(circles[i].radius), screenHeight-int32(circles[i].radius)))
				
				circles[i].position.Convert().X = px
				circles[i].position.Convert().Y = py
				circles[i].speed = float32(rl.GetRandomValue(1, 100)) / 2000.0
				circles[i].color = *colors[rl.GetRandomValue(0, 13)]
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		for i := (MAX_CIRCLES - 1); i >= 0; i-- {
			rl.DrawCircleV(circles[i].position, circles[i].radius, rl.Fade(circles[i].color, circles[i].alpha))
		}

		rl.DrawRectangle(20, screenHeight-20-12, screenWidth-40, 12, *rl.LightGray)
		rl.DrawRectangle(20, screenHeight-20-12, int32(timePlayed), 12, *rl.Maroon)
		rl.DrawRectangleLines(20, screenHeight-20-12, screenWidth-40, 12, *rl.Gray)

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(music)

	rl.CloseAudioDevice()
	
	rl.CloseWindow()
}

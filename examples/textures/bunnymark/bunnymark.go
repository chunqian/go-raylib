package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

var (
	MAX_BUNNIES        = 50000
	MAX_BATCH_ELEMENTS = 8192
)

type Bunny struct {
	position rl.Vector2
	speed    rl.Vector2
	color    rl.Color
}

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - bunnymark")

	texBunny := rl.LoadTexture("../resources/wabbit_alpha.png")
	bunnies := make([]Bunny, MAX_BUNNIES)

	bunniesCount := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {

			for i := 0; i < 100; i++ {

				if bunniesCount < MAX_BUNNIES {
					bunnies[bunniesCount].position = rl.GetMousePosition()
					bunnies[bunniesCount].speed.SetX(float32(rl.GetRandomValue(-250, 250)) / 60.0)
					bunnies[bunniesCount].speed.SetY(float32(rl.GetRandomValue(-250, 250)) / 60.0)
					bunnies[bunniesCount].color = rl.Color{
						R: byte(rl.GetRandomValue(50, 240)),
						G: byte(rl.GetRandomValue(80, 240)),
						B: byte(rl.GetRandomValue(100, 240)),
						A: 255,
					}
					bunniesCount++
				}
			}
		}

		for i := 0; i < bunniesCount; i++ {
			bunnies[i].position.SetX(*bunnies[i].position.GetX() + *bunnies[i].speed.GetX())
			bunnies[i].position.SetY(*bunnies[i].position.GetY() + *bunnies[i].speed.GetY())

			if (*bunnies[i].position.GetX()+float32(texBunny.Width)/2.0) > float32(rl.GetScreenWidth()) ||
				(*bunnies[i].position.GetX()+float32(texBunny.Width)/2.0) < 0.0 {
				bunnies[i].speed.SetX(*bunnies[i].speed.GetX() * (-1.0))
			}
			if (*bunnies[i].position.GetY()+float32(texBunny.Height)/2.0) > float32(rl.GetScreenHeight()) ||
				(*bunnies[i].position.GetY()+float32(texBunny.Height)/2.0-40.0) < 0.0 {
				bunnies[i].speed.SetY(*bunnies[i].speed.GetY() * (-1.0))
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		for i := 0; i < bunniesCount; i++ {
			rl.DrawTexture(texBunny, int32(*bunnies[i].position.GetX()), int32(*bunnies[i].position.GetY()), bunnies[i].color)
		}

		rl.DrawRectangle(0, 0, screenWidth, 40, *rl.Black)

		rl.DrawText(fmt.Sprintf("bunnies: %d", bunniesCount), 120, 10, 20, *rl.Green)

		rl.DrawText(fmt.Sprintf("batched draw calls: %d", 1+bunniesCount/MAX_BATCH_ELEMENTS), 320, 10, 20, *rl.Maroon)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(texBunny)

	rl.CloseWindow()
}

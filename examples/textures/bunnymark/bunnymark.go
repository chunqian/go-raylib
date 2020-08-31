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

	texBunny := rl.LoadTexture("../textures/resources/wabbit_alpha.png")

	bunnies := make([]Bunny, MAX_BUNNIES)
	for i := range bunnies {
		bunnies[i].position.PassRef()
		bunnies[i].speed.PassRef()
	}
	bunniesCount := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {

			for i := 0; i < 100; i++ {

				if bunniesCount < MAX_BUNNIES {
					bunnies[bunniesCount].position = rl.GetMousePosition()

					bunnies[bunniesCount].speed.This.X = float32(rl.GetRandomValue(-250, 250)) / 60.0
					bunnies[bunniesCount].speed.This.Y = float32(rl.GetRandomValue(-250, 250)) / 60.0

					bunnies[bunniesCount].color = rl.NewColor(
						byte(rl.GetRandomValue(50, 240)),
						byte(rl.GetRandomValue(80, 240)),
						byte(rl.GetRandomValue(100, 240)),
						255,
					)
					bunniesCount++
				}
			}
		}

		for i := 0; i < bunniesCount; i++ {

			bunnies[i].position.This.X += bunnies[i].speed.This.X
			bunnies[i].position.This.Y += bunnies[i].speed.This.Y

			if (bunnies[i].position.This.X+float32(texBunny.This.Width)/2.0) > float32(rl.GetScreenWidth()) ||
				(bunnies[i].position.This.X+float32(texBunny.This.Width)/2.0) < 0.0 {
				bunnies[i].speed.This.X *= -1.0
			}
			if (bunnies[i].position.This.Y+float32(texBunny.This.Height)/2.0) > float32(rl.GetScreenHeight()) ||
				(bunnies[i].position.This.Y+float32(texBunny.This.Height)/2.0-40.0) < 0.0 {
				bunnies[i].speed.This.Y *= -1.0
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := 0; i < bunniesCount; i++ {
			rl.DrawTexture(texBunny, int32(bunnies[i].position.This.X), int32(bunnies[i].position.This.Y), bunnies[i].color)
		}

		rl.DrawRectangle(0, 0, screenWidth, 40, rl.Black)

		rl.DrawText(fmt.Sprintf("bunnies: %d", bunniesCount), 120, 10, 20, rl.Green)

		rl.DrawText(fmt.Sprintf("batched draw calls: %d", 1+bunniesCount/MAX_BATCH_ELEMENTS), 320, 10, 20, rl.Maroon)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(texBunny)

	rl.CloseWindow()
}

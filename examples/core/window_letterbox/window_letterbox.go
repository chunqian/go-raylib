package main

import (
	rl "goray/raylib"

	"fmt"
	"math"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_WINDOW_RESIZABLE | rl.FLAG_VSYNC_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - window scale letterbox")
	defer rl.CloseWindow()

	rl.SetWindowMinSize(320, 240)

	gameScreenWidth := int32(640)
	gameScreenHeight := int32(480)

	target := rl.LoadRenderTexture(gameScreenWidth, gameScreenHeight)
	defer rl.UnloadRenderTexture(target)

	rl.SetTextureFilter(*target.Texture(), int32(rl.FILTER_BILINEAR))

	colors := [10]rl.Color{}
	for i := range colors {
		colors[i] = rl.NewColor(
			byte(rl.GetRandomValue(100, 250)),
			byte(rl.GetRandomValue(50, 150)),
			byte(rl.GetRandomValue(10, 100)),
			255,
		)
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		scale := float32(
			math.Min(
				float64(rl.GetScreenWidth())/float64(gameScreenWidth),
				float64(rl.GetScreenHeight())/float64(gameScreenHeight),
			),
		)

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {

			for i := range colors {
				colors[i] = rl.NewColor(
					byte(rl.GetRandomValue(100, 250)),
					byte(rl.GetRandomValue(50, 150)),
					byte(rl.GetRandomValue(10, 100)),
					255,
				)
			}
		}

		mouse := rl.GetMousePosition()
		virtualMouse := rl.NewVector2(
			(mouse.This.X-(float32(rl.GetScreenWidth())-(float32(gameScreenWidth)*scale))*0.5)/scale,
			(mouse.This.Y-(float32(rl.GetScreenHeight())-(float32(gameScreenHeight)*scale))*0.5)/scale,
		)
		virtualMouse = ClampValue(
			virtualMouse,
			rl.NewVector2(0, 0),
			rl.NewVector2(float32(gameScreenWidth), float32(gameScreenHeight)),
		)

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginTextureMode(target)

		rl.ClearBackground(rl.RayWhite) // Clear render texture background color

		for i := range colors {
			rl.DrawRectangle(
				0,
				(gameScreenHeight/10)*int32(i),
				gameScreenWidth,
				gameScreenHeight/10,
				colors[i],
			)
		}

		rl.DrawText("If executed inside a window,\nyou can resize the window,\nand see the screen scaling!", 10, 25, 20, rl.White)

		rl.DrawText(fmt.Sprintf("Default Mouse: [%d , %d]", int32(mouse.This.X), int32(mouse.This.Y)), 350, 25, 20, rl.Green)
		rl.DrawText(fmt.Sprintf("Virtual Mouse: [%d , %d]", int32(virtualMouse.This.X), int32(virtualMouse.This.Y)), 350, 55, 20, rl.Yellow)

		rl.EndTextureMode()

		rl.DrawTexturePro(
			*target.Texture(),
			rl.NewRectangle(0, 0, float32(target.Texture().This.Width), -float32(target.Texture().This.Height)),
			rl.NewRectangle(
				(float32(rl.GetScreenWidth())-(float32(gameScreenWidth)*scale))*0.5,
				(float32(rl.GetScreenHeight())-(float32(gameScreenHeight)*scale))*0.5,
				float32(gameScreenWidth)*scale,
				float32(gameScreenHeight)*scale,
			),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)

		rl.EndDrawing()
	}
}

func ClampValue(value rl.Vector2, min rl.Vector2, max rl.Vector2) rl.Vector2 {

	result := value
	if result.This.X > max.This.X {
		result.This.X = max.This.X
	}
	if result.This.X < min.This.X {
		result.This.X = min.This.X
	}
	if result.This.Y > max.This.Y {
		result.This.Y = max.This.Y
	}
	if result.This.Y < min.This.Y {
		result.This.Y = min.This.Y
	}
	return result
}

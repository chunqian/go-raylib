package main

import (
	rl "go-raylib/raylib"
	"runtime"
)

var MAX_BUILDINGS = 100

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 2d camera")
	defer rl.CloseWindow()

	player := rl.NewRectangle(400, 280, 40, 40)

	buildings := make([]rl.Rectangle, MAX_BUILDINGS)
	buildColors := make([]rl.Color, MAX_BUILDINGS)

	spacing := int32(0)

	for i := 0; i < MAX_BUILDINGS; i++ {
		buildings[i].Width = float32(rl.GetRandomValue(50, 200))
		buildings[i].Height = float32(rl.GetRandomValue(100, 800))
		buildings[i].Y = float32(screenHeight) - 130 - buildings[i].Height
		buildings[i].X = float32(-6000 + spacing)

		spacing += int32(buildings[i].Width)

		buildColors[i] = rl.NewColor(
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 250)),
			255,
		)
	}

	camera := rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(player.X+20, player.Y+20),
		0.0,
		1.0,
	)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			player.X += 2
		} else if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			player.X -= 2
		}

		camera.Target = rl.NewVector2(player.X+20, player.Y+20)

		if rl.IsKeyDown(int32(rl.KEY_A)) {
			camera.Rotation--
		} else if rl.IsKeyDown(int32(rl.KEY_S)) {
			camera.Rotation++
		}

		if camera.Rotation > 40 {
			camera.Rotation = 40
		} else if camera.Rotation < -40 {
			camera.Rotation = -40
		}

		camera.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if camera.Zoom > 3.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom < 0.1 {
			camera.Zoom = 0.1
		}

		if rl.IsKeyPressed(int32(rl.KEY_R)) {
			camera.Zoom = 1.0
			camera.Rotation = 0.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(-6000, 320, 13000, 8000, rl.DarkGray)

		for i := 0; i < MAX_BUILDINGS; i++ {
			rl.DrawRectangleRec(buildings[i], buildColors[i])
		}

		rl.DrawRectangleRec(player, rl.Red)

		rl.DrawLine(
			int32(camera.Target.X),
			-screenHeight*10,
			int32(camera.Target.X),
			screenHeight*10, rl.Green,
		)
		rl.DrawLine(
			-screenWidth*10,
			int32(camera.Target.Y),
			screenWidth*10,
			int32(camera.Target.Y),
			rl.Green,
		)

		rl.EndMode2D()

		rl.DrawText("SCREEN AREA", 640, 10, 20, rl.Red)

		rl.DrawRectangle(0, 0, screenWidth, 5, rl.Red)
		rl.DrawRectangle(0, 5, 5, screenHeight-10, rl.Red)
		rl.DrawRectangle(screenWidth-5, 5, 5, screenHeight-10, rl.Red)
		rl.DrawRectangle(0, screenHeight-5, screenWidth, 5, rl.Red)

		rl.DrawRectangle(10, 10, 250, 113, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 250, 113, rl.Blue)

		rl.DrawText("Free 2d camera controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Right/Left to move Offset", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- A / S to Rotate", 40, 80, 10, rl.DarkGray)
		rl.DrawText("- R to reset Zoom and Rotation", 40, 100, 10, rl.DarkGray)

		rl.EndDrawing()
	}
}

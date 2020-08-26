package main

import (
	rl "goray/raylib"
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

	player := rl.NewRectangle(400, 280, 40, 40)
	playerT := player.Convert()

	buildings := make([]rl.Rectangle, MAX_BUILDINGS)
	buildColors := make([]rl.Color, MAX_BUILDINGS)

	spacing := int32(0)

	for i := 0; i < MAX_BUILDINGS; i++ {
		buildingT := buildings[i].Convert()
		buildingT.Width = float32(rl.GetRandomValue(50, 200))
		buildingT.Height = float32(rl.GetRandomValue(100, 800))
		buildingT.Y = float32(screenHeight) - 130 - buildingT.Height
		buildingT.X = float32(-6000 + spacing)

		spacing += int32(buildingT.Width)

		buildColors[i] = rl.NewColor(
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 240)),
			byte(rl.GetRandomValue(200, 250)),
			255,
		)
	}

	camera := rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(playerT.X+20, playerT.Y+20),
		0.0,
		1.0,
	)
	cameraT := camera.Convert()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(int32(rl.KEY_RIGHT)) {
			playerT.X += 2
		} else if rl.IsKeyDown(int32(rl.KEY_LEFT)) {
			playerT.X -= 2
		}

		cameraT.Target, _ = rl.NewVector2(playerT.X+20, playerT.Y+20).PassValue()

		if rl.IsKeyDown(int32(rl.KEY_A)) {
			cameraT.Rotation--
		} else if rl.IsKeyDown(int32(rl.KEY_S)) {
			cameraT.Rotation++
		}

		if cameraT.Rotation > 40 {
			cameraT.Rotation = 40
		} else if cameraT.Rotation < -40 {
			cameraT.Rotation = -40
		}

		cameraT.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if cameraT.Zoom > 3.0 {
			cameraT.Zoom = 3.0
		} else if cameraT.Zoom < 0.1 {
			cameraT.Zoom = 0.1
		}

		if rl.IsKeyPressed(int32(rl.KEY_R)) {
			cameraT.Zoom = 1.0
			cameraT.Rotation = 0.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangle(-6000, 320, 13000, 8000, rl.DarkGray)

		for i := 0; i < MAX_BUILDINGS; i++ {
			rl.DrawRectangleRec(buildings[i], buildColors[i])
		}

		rl.DrawRectangleRec(player, rl.Red)

		targetT := camera.GetTarget().Convert()
		rl.DrawLine(int32(targetT.X), -screenHeight*10, int32(targetT.X), screenHeight*10, rl.Green)
		rl.DrawLine(-screenWidth*10, int32(targetT.Y), screenWidth*10, int32(targetT.Y), rl.Green)

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

	rl.CloseWindow()
}

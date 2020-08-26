package main

import (
	rl "goray/raylib"

	"runtime"
)

var (
	MAX_COLUMNS = 20
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera first person")

	camera := rl.NewCamera(
		rl.NewVector3(4.0, 2.0, 4.0),
		rl.NewVector3(0.0, 1.8, 0.0),
		rl.NewVector3(0.0, 1.0, 0.0),
		60.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	heights := make([]float32, MAX_COLUMNS)
	positions := make([]rl.Vector3, MAX_COLUMNS)
	colors := make([]rl.Color, MAX_COLUMNS)

	for i := 0; i < MAX_COLUMNS; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = rl.NewVector3(
			float32(rl.GetRandomValue(-15, 15)),
			heights[i]/2,
			float32(rl.GetRandomValue(-15, 15)),
		)
		colors[i] = rl.NewColor(
			byte(rl.GetRandomValue(20, 255)),
			byte(rl.GetRandomValue(10, 55)),
			30,
			255,
		)
	}

	rl.SetCameraMode(camera, int32(rl.CAMERA_FIRST_PERSON))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawPlane(
			rl.NewVector3(0.0, 0.0, 0.0),
			rl.NewVector2(32.0, 32.0),
			rl.LightGray,
		)
		rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)
		rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)
		rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)

		for i := 0; i < MAX_COLUMNS; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawText("First person camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Move with keys: W, A, S, D", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse move to look around", 40, 60, 10, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

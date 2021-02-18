package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"runtime"
)

const (
	FOVY_PERSPECTIVE   = 45.0
	WIDTH_ORTHOGRAPHIC = 10.0
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - geometric shapes")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(0, 10.0, 10.0),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		FOVY_PERSPECTIVE,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			if camera.Type == int32(rl.CAMERA_PERSPECTIVE) {
				camera.Fovy = WIDTH_ORTHOGRAPHIC
				camera.Type = int32(rl.CAMERA_ORTHOGRAPHIC)
			} else {
				camera.Fovy = FOVY_PERSPECTIVE
				camera.Type = int32(rl.CAMERA_PERSPECTIVE)
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawCube(rl.NewVector3(-4, 0, 2), 2, 5, 2, rl.Red)
		rl.DrawCubeWires(rl.NewVector3(-4, 0, 2), 2, 5, 2, rl.Gold)
		rl.DrawCubeWires(rl.NewVector3(-4, 0, -2), 3, 6, 2, rl.Maroon)

		rl.DrawSphere(rl.NewVector3(-1, 0, -2), 1, rl.Green)
		rl.DrawSphereWires(rl.NewVector3(1, 0, 2), 2, 16, 16, rl.Lime)

		rl.DrawCylinder(rl.NewVector3(4, 0, -2), 1, 2, 3, 4, rl.SkyBlue)
		rl.DrawCylinderWires(rl.NewVector3(4, 0, -2), 1, 2, 3, 4, rl.DarkBlue)
		rl.DrawCylinderWires(rl.NewVector3(4.5, -1, 2), 1, 1, 2, 6, rl.Brown)

		rl.DrawCylinder(rl.NewVector3(1, 0, -4), 0, 1.5, 3, 8, rl.Gold)
		rl.DrawCylinderWires(rl.NewVector3(1, 0, -4), 0, 1.5, 3, 8, rl.Pink)

		rl.DrawGrid(10, 1)

		rl.EndMode3D()

		rl.DrawText("Press Spacebar to switch camera type", 10, rl.GetScreenHeight()-30, 20, rl.DarkGray)

		if camera.Type == int32(rl.CAMERA_ORTHOGRAPHIC) {
			rl.DrawText("ORTHOGRAPHIC", 10, 40, 20, rl.Black)
		} else if camera.Type == int32(rl.CAMERA_PERSPECTIVE) {
			rl.DrawText("PERSPECTIVE", 10, 40, 20, rl.Black)
		}

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera free")

	camera := rl.NewCamera(
		rl.NewVector3(10.0, 10.0, 10.0),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	cubePosition := rl.NewVector3(0, 0, 0)
	cubeScreenPosition := rl.NewVector2(0, 0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_FREE))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		cubeScreenPosition = rl.GetWorldToScreen(
			rl.NewVector3(cubePosition.X, cubePosition.Y+2.5, cubePosition.Z),
			camera,
		)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText(
			"Enemy: 100 / 100",
			int32(cubeScreenPosition.X)-rl.MeasureText("Enemy: 100/100", 20)/2,
			int32(cubeScreenPosition.Y),
			20,
			rl.Black,
		)
		rl.DrawText(
			"Text is always on top of the cube",
			(screenWidth-rl.MeasureText("Text is always on top of the cube", 20))/2,
			25,
			20,
			rl.Gray,
		)

		rl.EndDrawing()
	}
}

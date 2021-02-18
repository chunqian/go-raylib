package main

import (
	rl "go-raylib/raylib"

	"runtime"
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
		rl.NewVector3(0, 10, 10),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawCube(rl.NewVector3(-4.0, 0, 2.0), 2.0, 5.0, 2.0, rl.Red)
		rl.DrawCubeWires(rl.NewVector3(-4.0, 0, 2.0), 2.0, 5.0, 2.0, rl.Gold)
		rl.DrawCubeWires(rl.NewVector3(-4.0, 0, -2.0), 3.0, 6.0, 2.0, rl.Maroon)

		rl.DrawSphere(rl.NewVector3(-1.0, 0, -2.0), 1.0, rl.Green)
		rl.DrawSphereWires(rl.NewVector3(1.0, 0, -2.0), 2.0, 16, 16, rl.Lime)

		rl.DrawCylinder(rl.NewVector3(4.0, 0, -2.0), 1.0, 2.0, 3.0, 4, rl.SkyBlue)
		rl.DrawCylinderWires(rl.NewVector3(4.0, 0, -2.0), 1.0, 2.0, 3.0, 4, rl.DarkBlue)
		rl.DrawCylinderWires(rl.NewVector3(1.0, 0, -4.0), 0, 1.5, 3.0, 8, rl.Pink)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

package main

import (
	rl "goray/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - drawing billboards")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(5.0, 4.0, 5.0),
		rl.NewVector3(0, 2.0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	bill := rl.LoadTexture("../models/resources/billboard.png")
	defer rl.UnloadTexture(bill)
	billPosition := rl.NewVector3(0, 2.0, 0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawGrid(10, 1.0)

		rl.DrawBillboard(camera, bill, billPosition, 2.0, rl.White)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

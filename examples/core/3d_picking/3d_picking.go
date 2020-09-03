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

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d picking")
	defer rl.CloseWindow()

	camera := rl.NewCamera3D(
		rl.NewVector3(10, 10, 10),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	cubePosition := rl.NewVector3(0, 1.0, 0)
	cubeSize := rl.NewVector3(2.0, 2.0, 2.0)

	var ray rl.Ray

	collision := false

	rl.SetCameraMode(rl.Camera(camera), int32(rl.CAMERA_FREE))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera((*rl.Camera)(&camera))

		if rl.IsMouseButtonPressed(int32(rl.MOUSE_LEFT_BUTTON)) {
			if !collision {
				ray = rl.GetMouseRay(rl.GetMousePosition(), rl.Camera(camera))

				collision = rl.CheckCollisionRayBox(
					ray,
					rl.NewBoundingBox(
						rl.NewVector3(
							cubePosition.This.X-cubeSize.This.X/2,
							cubePosition.This.Y-cubeSize.This.Y/2,
							cubePosition.This.Z-cubeSize.This.Z/2,
						),
						rl.NewVector3(
							cubePosition.This.X+cubeSize.This.X/2,
							cubePosition.This.Y+cubeSize.This.Y/2,
							cubePosition.This.Z+cubeSize.This.Z/2,
						),
					),
				)
			} else {
				collision = false
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		if collision {
			rl.DrawCube(cubePosition, cubeSize.This.X, cubeSize.This.Y, cubeSize.This.Z, rl.Red)
			rl.DrawCubeWires(cubePosition, cubeSize.This.X, cubeSize.This.Y, cubeSize.This.Z, rl.Maroon)

			rl.DrawCubeWires(cubePosition, cubeSize.This.X+0.2, cubeSize.This.Y+0.2, cubeSize.This.Z+0.2, rl.Green)
		} else {
			rl.DrawCube(cubePosition, cubeSize.This.X, cubeSize.This.Y, cubeSize.This.Z, rl.Gray)
			rl.DrawCubeWires(cubePosition, cubeSize.This.X, cubeSize.This.Y, cubeSize.This.Z, rl.DarkGray)
		}

		rl.DrawRay(ray, rl.Maroon)
		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText("Try selecting the box with mouse!", 240, 10, 20, rl.DarkGray)

		if collision {
			rl.DrawText(
				"BOX SELECTED",
				(screenWidth-rl.MeasureText("BOX SELECTED", 30))/2,
				int32(float32(screenHeight)*0.1),
				30,
				rl.Green,
			)
		}

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

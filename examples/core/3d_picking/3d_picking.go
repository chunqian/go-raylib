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
							cubePosition.Convert().X-cubeSize.Convert().X/2,
							cubePosition.Convert().Y-cubeSize.Convert().Y/2,
							cubePosition.Convert().Z-cubeSize.Convert().Z/2,
						),
						rl.NewVector3(
							cubePosition.Convert().X+cubeSize.Convert().X/2,
							cubePosition.Convert().Y+cubeSize.Convert().Y/2,
							cubePosition.Convert().Z+cubeSize.Convert().Z/2,
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

		cubeSizeT := cubeSize.Convert()
		if collision {
			rl.DrawCube(cubePosition, cubeSizeT.X, cubeSizeT.Y, cubeSizeT.Z, rl.Red)
			rl.DrawCubeWires(cubePosition, cubeSizeT.X, cubeSizeT.Y, cubeSizeT.Z, rl.Maroon)

			rl.DrawCubeWires(cubePosition, cubeSizeT.X+0.2, cubeSizeT.Y+0.2, cubeSizeT.Z+0.2, rl.Green)
		} else {
			rl.DrawCube(cubePosition, cubeSizeT.X, cubeSizeT.Y, cubeSizeT.Z, rl.Gray)
			rl.DrawCubeWires(cubePosition, cubeSizeT.X, cubeSizeT.Y, cubeSizeT.Z, rl.DarkGray)
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

	rl.CloseWindow()
}

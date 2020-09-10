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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - first person maze")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(0.2, 0.4, 0.2),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	imMap := rl.LoadImage("../models/resources/cubicmap.png")
	cubicmap := rl.LoadTextureFromImage(imMap)
	defer rl.UnloadTexture(cubicmap)

	mesh := rl.GenMeshCubicmap(imMap, rl.NewVector3(1.0, 1.0, 1.0))
	model := rl.LoadModelFromMesh(mesh)
	defer rl.UnloadModel(model)

	texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
	defer rl.UnloadTexture(texture)

	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture

	mapPixels := rl.GetImageData(imMap)
	defer rl.UnloadColors(mapPixels)

	rl.UnloadImage(imMap)

	mapPosition := rl.NewVector3(-16, 0, -8)
	// playerPosition := camera.Position

	rl.SetCameraMode(camera, int32(rl.CAMERA_FIRST_PERSON))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		oldCamPos := camera.Position

		rl.UpdateCamera(&camera)

		playerPos := rl.NewVector2(camera.Position.X, camera.Position.Z)
		playerRadius := float32(0.1)

		playerCellX := int32(playerPos.X - mapPosition.X + 0.5)
		playerCellY := int32(playerPos.Y - mapPosition.Z + 0.5)

		if playerCellX < 0 {
			playerCellX = 0
		} else if playerCellX >= cubicmap.Width {
			playerCellX = cubicmap.Width - 1
		}

		if playerCellY < 0 {
			playerCellY = 0
		} else if playerCellY >= cubicmap.Height {
			playerCellY = cubicmap.Height - 1
		}

		for y := int32(0); y < cubicmap.Height; y++ {

			for x := int32(0); x < cubicmap.Width; x++ {

				if mapPixels.Index(y*cubicmap.Width+x).R == 255 &&
					rl.CheckCollisionCircleRec(
						playerPos,
						playerRadius,
						rl.NewRectangle(
							mapPosition.X-0.5+float32(x)*1.0,
							mapPosition.Z-0.5+float32(y)*1.0,
							1.0,
							1.0,
						),
					) {
					camera.Position = oldCamPos
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, mapPosition, 1.0, rl.White)

		rl.EndMode3D()

		rl.DrawTextureEx(
			cubicmap,
			rl.NewVector2(
				float32(rl.GetScreenWidth()-cubicmap.Width*4-20),
				20,
			),
			0,
			4,
			rl.White,
		)
		rl.DrawRectangleLines(
			rl.GetScreenWidth()-cubicmap.Width*4-20,
			20,
			cubicmap.Width*4,
			cubicmap.Height*4,
			rl.Green,
		)

		rl.DrawRectangle(
			rl.GetScreenWidth()-cubicmap.Width*4-20+playerCellX*4,
			20+playerCellY*4,
			4,
			4,
			rl.Red,
		)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

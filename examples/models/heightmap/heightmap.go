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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - heightmap loading and drawing")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(18, 18, 18),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	image := rl.LoadImage("../models/resources/heightmap.png")
	texture := rl.LoadTextureFromImage(image)
	defer rl.UnloadTexture(texture)

	mesh := rl.GenMeshHeightmap(image, rl.NewVector3(16, 8, 16))
	model := rl.LoadModelFromMesh(mesh)
	defer rl.UnloadModel(model)

	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture
	mapPosition := rl.NewVector3(-8, 0, -8)

	rl.UnloadImage(image)

	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, mapPosition, 1.0, rl.Red)

		rl.DrawGrid(20, 1.0)

		rl.EndMode3D()

		rl.DrawTexture(texture, screenWidth-texture.Width-20, 20, rl.White)
		rl.DrawRectangleLines(screenWidth-texture.Width-20, 20, texture.Width, texture.Height, rl.Green)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

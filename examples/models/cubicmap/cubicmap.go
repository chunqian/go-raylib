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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - cubesmap loading and drawing")

	camera := rl.NewCamera(
		rl.NewVector3(16, 14, 16),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45,
		int32(rl.CAMERA_CUSTOM),
	)

	image := rl.LoadImage("../models/resources/cubicmap.png")
	cubicmap := rl.LoadTextureFromImage(image)

	mesh := rl.GenMeshCubicmap(image, rl.NewVector3(1, 1, 1))
	model := rl.LoadModelFromMesh(mesh)

	texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
	model.Materials(0).Maps(rl.MAP_DIFFUSE).This.Texture, _ = texture.PassValue()

	rl.UnloadImage(image)
	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.NewVector3(-16, 0, -8), 1.0, rl.White)

		rl.EndMode3D()

		rl.DrawTextureEx(cubicmap,
			rl.NewVector2(float32(screenWidth-cubicmap.This.Width*4-20), 20),
			0, 4.0,
			rl.White,
		)
		rl.DrawRectangleLines(screenWidth-cubicmap.This.Width*4-20, 20, cubicmap.This.Width*4, cubicmap.This.Height*4, rl.Green)

		rl.DrawText("cubicmap image used to", 658, 90, 10, rl.Gray)
		rl.DrawText("generate map 3d model", 658, 104, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(cubicmap)
	rl.UnloadTexture(texture)
	rl.UnloadModel(model)

	rl.CloseWindow()
}

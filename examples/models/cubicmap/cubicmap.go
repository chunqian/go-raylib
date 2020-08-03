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

	camera := rl.Camera{
		Position: rl.Vector3{X: 16.0, Y: 14.0, Z: 16.0},
		Target:   rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Up:       rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:     45.0,
		Type:     int32(rl.CAMERA_CUSTOM),
	}

	image := rl.LoadImage("../models/resources/cubicmap.png")
	cubicmap := rl.LoadTextureFromImage(image)

	mesh := rl.GenMeshCubicmap(image, rl.Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	model := rl.LoadModelFromMesh(mesh)

	texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
	model.GetMaterials(0).GetMaps(rl.MAP_DIFFUSE).Convert().Texture = *texture.Ref()

	rl.UnloadImage(image)
	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.Vector3{X: -16.0, Y: 0.0, Z: -8.0}, 1.0, *rl.White)

		rl.EndMode3D()

		rl.DrawTextureEx(cubicmap, rl.Vector2{X: float32(screenWidth - cubicmap.Width*4 - 20), Y: 20}, 0, 4.0, *rl.White)
		rl.DrawRectangleLines(screenWidth-cubicmap.Width*4-20, 20, cubicmap.Width*4, cubicmap.Height*4, *rl.Green)

		rl.DrawText("cubicmap image used to", 658, 90, 10, *rl.Gray)
		rl.DrawText("generate map 3d model", 658, 104, 10, *rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadTexture(cubicmap)
	rl.UnloadTexture(texture)
	rl.UnloadModel(model)

	rl.CloseWindow()
}

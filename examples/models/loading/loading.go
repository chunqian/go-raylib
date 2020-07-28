package main

import (
	"rl"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading")

	camera := rl.Camera{
		Position: rl.Vector3{X: 50.0, Y: 50.0, Z: 50.0},
		Target:   rl.Vector3{X: 0.0, Y: 10.0, Z: 0.0},
		Up:       rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:     45.0,
		Type:     int32(rl.CAMERA_PERSPECTIVE),
	}

	model := rl.LoadModel("../resources/models/castle.obj")
	texture := rl.LoadTexture("../resources/models/castle_diffuse.png")

	model.GetMaterials(*model.GetMaterialCount())[0].GetMaps(rl.MAP_DIFFUSE + 1)[rl.MAP_DIFFUSE].SetTexture(&texture)

	bounds := rl.MeshBoundingBox(model.GetMeshes(1)[0])

	rl.SetCameraMode(camera, int32(rl.CAMERA_FREE))

	selected := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsFileDropped() {
			count := int32(0)
			droppedFiles := rl.GetDroppedFiles(&count)

			if count == 1 {
				if rl.IsFileExtension(string(droppedFiles[0][:]), ".obj") ||
					rl.IsFileExtension(string(droppedFiles[0][:]), ".gltf") ||
					rl.IsFileExtension(string(droppedFiles[0][:]), ".iqm") {
					rl.UnloadModel(model)
					model = rl.LoadModel(string(droppedFiles[0][:]))
					model.GetMaterials(*model.GetMaterialCount())[0].GetMaps(rl.MAP_DIFFUSE + 1)[rl.MAP_DIFFUSE].SetTexture(&texture)
					bounds = rl.MeshBoundingBox(model.GetMeshes(1)[0])
				} else if rl.IsFileExtension(string(droppedFiles[0][:]), ".png") {

					rl.UnloadTexture(texture)
					texture = rl.LoadTexture(string(droppedFiles[0][:]))
					model.GetMaterials(*model.GetMaterialCount())[0].GetMaps(rl.MAP_DIFFUSE + 1)[rl.MAP_DIFFUSE].SetTexture(&texture)
				}
			}
			rl.ClearDroppedFiles()
		}

		if rl.IsMouseButtonPressed(int32(rl.MOUSE_LEFT_BUTTON)) {
			if rl.CheckCollisionRayBox(rl.GetMouseRay(rl.GetMousePosition(), camera), bounds) {
				selected = !selected
			} else {
				selected = false
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(*rl.RayWhite)
		rl.BeginMode3D(rl.Camera3D(camera))
		rl.DrawModel(model, rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 1.0, *rl.White)
		rl.DrawGrid(20, 10.0)
		if selected {
			rl.DrawBoundingBox(bounds, *rl.Green)
		}
		rl.EndMode3D()
		rl.DrawText("Drag & drop model to load mesh/texture.", 10, rl.GetScreenHeight()-20, 10, *rl.DarkGray)
		if selected {
			rl.DrawText("MODEL SELECTED", rl.GetScreenWidth()-110, 10, 10, *rl.Green)
		}
		rl.DrawText("(c) Castle 3D model by Alberto Cano", screenWidth-200, screenHeight-20, 10, *rl.Gray)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
	rl.UnloadModel(model)
	rl.CloseWindow()
}

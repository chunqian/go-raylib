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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading")

	camera := rl.NewCamera(
		rl.NewVector3(50, 50, 50),
		rl.NewVector3(0, 10, 0),
		rl.NewVector3(0, 1.0, 0),
		45,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	model := rl.LoadModel("../models/resources/models/castle.obj")
	texture := rl.LoadTexture("../models/resources/models/castle_diffuse.png")

	model.GetMaterials(0).GetMaps(rl.MAP_DIFFUSE).Convert().Texture, _ = texture.PassValue()

	bounds := rl.MeshBoundingBox(*model.GetMeshes(0))

	rl.SetCameraMode(camera, int32(rl.CAMERA_FREE))

	selected := false

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsFileDropped() {
			count := int32(0)
			droppedFiles := rl.GetDroppedFiles(&count)
			droppedFilePath := rl.StringFromPPByte(droppedFiles, 0)

			if count == 1 {
				if rl.IsFileExtension(droppedFilePath, ".obj") ||
					rl.IsFileExtension(droppedFilePath, ".gltf") ||
					rl.IsFileExtension(droppedFilePath, ".iqm") {
					rl.UnloadModel(model)
					model = rl.LoadModel(droppedFilePath)
					model.GetMaterials(0).GetMaps(rl.MAP_DIFFUSE).Convert().Texture, _ = texture.PassValue()
					bounds = rl.MeshBoundingBox(*model.GetMeshes(0))
				} else if rl.IsFileExtension(droppedFilePath, ".png") {

					rl.UnloadTexture(texture)
					texture = rl.LoadTexture(droppedFilePath)
					model.GetMaterials(0).GetMaps(rl.MAP_DIFFUSE).Convert().Texture, _ = texture.PassValue()
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
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(rl.Camera3D(camera))
		rl.DrawModel(model, rl.NewVector3(0, 0, 0), 1.0, rl.White)
		rl.DrawGrid(20, 10.0)
		if selected {
			rl.DrawBoundingBox(bounds, rl.Green)
		}
		rl.EndMode3D()
		rl.DrawText("Drag & drop model to load mesh/texture.", 10, rl.GetScreenHeight()-20, 10, rl.DarkGray)
		if selected {
			rl.DrawText("MODEL SELECTED", rl.GetScreenWidth()-110, 10, 10, rl.Green)
		}
		rl.DrawText("(c) Castle 3D model by Alberto Cano", screenWidth-200, screenHeight-20, 10, rl.Gray)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
	rl.UnloadModel(model)
	rl.CloseWindow()
}

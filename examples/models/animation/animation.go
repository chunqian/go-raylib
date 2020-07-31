package main

import (
	rl "goray/raylib"

	"runtime"
	// "runtime/debug"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - model animation")

	camera := rl.Camera{
		Position: rl.Vector3{X: 10.0, Y: 10.0, Z: 10.0},
		Target:   rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Up:       rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:     45.0,
		Type:     int32(rl.CAMERA_PERSPECTIVE),
	}

	// debug.PrintStack()
	model := rl.LoadModel("../resources/guy/guy.iqm")
	texture := rl.LoadTexture("../resources/guy/guytex.png")

	rl.SetMaterialTexture(&(model.GetMaterials(*model.GetMaterialCount())[0]), rl.MAP_DIFFUSE, texture)

	animsCount := int32(0)
	anims := rl.LoadModelAnimations("../resources/guy/guyanim.iqm", &animsCount)
	animFrameCounter := int32(0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_CUSTOM))
	rl.SetTargetFPS(60)

	framePoses := anims[0].GetFramePoses(*anims[0].GetFrameCount(), *model.GetBoneCount())

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsKeyDown(int32(rl.KEY_SPACE)) {
			animFrameCounter++
			rl.UpdateModelAnimation(model, anims[0], animFrameCounter)
			if animFrameCounter >= *anims[0].GetFrameCount() {
				animFrameCounter = 0
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(*rl.RayWhite)
		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModelEx(model, rl.Vector3{X: 0.0, Y: -5.0, Z: 0.0}, rl.Vector3{X: 1.0, Y: 0.0, Z: 0.0}, -90.0, rl.Vector3{X: 1.0, Y: 1.0, Z: 1.0}, *rl.White)

		for i := int32(0); i < *model.GetBoneCount(); i++ {
			rl.DrawCube(*(framePoses[animFrameCounter][i].GetTranslation()), 0.2, 0.2, 0.2, *rl.Red)
		}

		rl.DrawGrid(10, 1.0)
		rl.EndMode3D()

		rl.DrawText("PRESS SPACE to PLAY MODEL ANIMATION", 10, 10, 20, *rl.Maroon)
		rl.DrawText("(c) Guy IQM 3D model by @culacant", screenWidth-200, screenHeight-20, 10, *rl.Gray)

		rl.DrawFPS(10, 30)
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
	for i := int32(0); i < animsCount; i++ {
		rl.UnloadModelAnimation(anims[i])
	}

	rl.UnloadModel(model)
	rl.CloseWindow()
}

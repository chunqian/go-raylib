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
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(10, 10, 10),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	// debug.PrintStack()
	model := rl.LoadModel("../models/resources/guy/guy.iqm")
	defer rl.UnloadModel(model)

	texture := rl.LoadTexture("../models/resources/guy/guytex.png")
	defer rl.UnloadTexture(texture)

	rl.SetMaterialTexture(model.Materials(0), rl.MAP_DIFFUSE, texture)

	animsCount := int32(0)
	anims := rl.LoadModelAnimations("../models/resources/guy/guyanim.iqm", &animsCount)
	defer func() {
		for i := int32(0); i < animsCount; i++ {
			rl.UnloadModelAnimation(*anims.Index(i))
		}
	}()

	animFrameCounter := int32(0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_CUSTOM))
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		if rl.IsKeyDown(int32(rl.KEY_SPACE)) {
			animFrameCounter++
			rl.UpdateModelAnimation(model, *anims.Index(0), animFrameCounter)
			if animFrameCounter >= anims.Index(0).FrameCount {
				animFrameCounter = 0
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(rl.ToCamera3D(camera))

		rl.DrawModelEx(model, rl.NewVector3(0, -5, 0), rl.NewVector3(1, 0, 0), -90, rl.NewVector3(1, 1, 1), rl.White)

		for i := int32(0); i < model.BoneCount; i++ {
			framePose := anims.Index(0).FramePoses(animFrameCounter, i)
			rl.DrawCube(framePose.Translation, 0.2, 0.2, 0.2, rl.Red)
		}

		rl.DrawGrid(10, 1.0)
		rl.EndMode3D()

		rl.DrawText("PRESS SPACE to PLAY MODEL ANIMATION", 10, 10, 20, rl.Maroon)
		rl.DrawText("(c) Guy IQM 3D model by @culacant", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.DrawFPS(10, 30)
		rl.EndDrawing()
	}
}

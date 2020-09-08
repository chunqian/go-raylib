package main

import (
	rl "goray/raylib"

	"runtime"
)

var (
	CUBEMAP_SIZE = 512
	IRRADIANCE_SIZE = 32
	PREFILTERED_SIZE = 256
	BRDF_SIZE = 512
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))
	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - pbr material")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(4.0, 4.0, 4.0),
		rl.NewVector3(0, 0.5, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	model := rl.LoadModel("../models/resources/pbr/trooper.obj")
	defer rl.UnloadModel(model)

	rl.MeshTangents(model.Meshes(0))

	rl.UnloadMaterial(*model.Materials(0))
}

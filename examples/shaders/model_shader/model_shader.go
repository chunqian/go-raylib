package main

import (
	rl "goray/raylib"

	"fmt"
	"runtime"
)

const GLSL_VERSION = 330

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(uint32(rl.FLAG_MSAA_4X_HINT))

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - model shader")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(4.0, 4.0, 4.0),
		rl.NewVector3(0, 1.0, -1.0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	model := rl.LoadModel("../shaders/resources/models/watermill.obj")
	texture := rl.LoadTexture("../shaders/resources/models/watermill_diffuse.png")

	shader := rl.LoadShader("", fmt.Sprintf("../shaders/resources/shaders/glsl%d/grayscale.fs", GLSL_VERSION))
	defer func() {
		rl.UnloadShader(shader)
		rl.UnloadTexture(texture)
		rl.UnloadModel(model)
	}()

	model.Materialser(0).Shader = shader
	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = texture

	position := rl.NewVector3(0, 0, 0)

	rl.SetCameraMode(camera, int32(rl.CAMERA_FREE))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, position, 0.2, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawText("(c) Watermill 3D model by Alberto Cano", screenWidth-210, screenHeight-20, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}

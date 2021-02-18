package main

import (
	rl "go-raylib/raylib"

	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - skybox loading and drawing")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(1.0, 1.0, 1.0),
		rl.NewVector3(4.0, 1.0, 4.0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	cube := rl.GenMeshCube(1.0, 1.0, 1.0)
	skybox := rl.LoadModelFromMesh(cube)
	defer func() {
		rl.UnloadTexture(
			rl.Texture2D(skybox.Materialser(0).Mapser(int32(rl.MAP_CUBEMAP)).Texture),
		)
		rl.UnloadShader(skybox.Materialser(0).Shader)

		rl.UnloadModel(skybox)
	}()

	skybox.Materialser(0).Shader = rl.LoadShader("../models/resources/shaders/glsl330/skybox.vs", "../models/resources/shaders/glsl330/skybox.fs")
	rl.SetShaderValue(
		skybox.Materialser(0).Shader,
		rl.GetShaderLocation(skybox.Materialser(0).Shader, "environmentMap"),
		unsafe.Pointer(&[1]int32{int32(rl.MAP_CUBEMAP)}),
		int32(rl.UNIFORM_INT),
	)
	rl.SetShaderValue(
		skybox.Materialser(0).Shader,
		rl.GetShaderLocation(skybox.Materialser(0).Shader, "vflipped"),
		unsafe.Pointer(&[1]int32{1}),
		int32(rl.UNIFORM_INT),
	)

	shdrCubemap := rl.LoadShader("../models/resources/shaders/glsl330/cubemap.vs", "../models/resources/shaders/glsl330/cubemap.fs")
	rl.SetShaderValue(
		shdrCubemap,
		rl.GetShaderLocation(shdrCubemap, "equirectangularMap"),
		unsafe.Pointer(&[1]int32{0}),
		int32(rl.UNIFORM_INT),
	)

	texHDR := rl.LoadTexture("../models/resources/dresden_square.hdr")
	skybox.Materialser(0).Mapser(int32(rl.MAP_CUBEMAP)).Texture = rl.Texture(
		rl.GenTextureCubemap(shdrCubemap, texHDR, 512, int32(rl.UNCOMPRESSED_R8G8B8A8)),
	)

	rl.UnloadTexture(texHDR)
	rl.UnloadShader(shdrCubemap)

	rl.SetCameraMode(camera, int32(rl.CAMERA_FIRST_PERSON))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(skybox, rl.NewVector3(0, 0, 0), 1.0, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
